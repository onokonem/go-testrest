package main

import (
	"fmt"
	"net/http"

	reform "gopkg.in/reform.v1"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/onokonem/go-testrest/proto/restapi/operations/account"
)

//go:generate reform

//reform:account
type Account struct {
	ID     int64   `json:"id,omitempty" reform:"id,pk"`
	Name   string  `json:"name,omitempty" reform:"name"`
	Amount int64   `json:"amount,omitempty" reform:"amount"`
	Ctime  *string `json:"ctime,omitempty" reform:"ctime"`
	Mtime  *string `json:"mtime,omitempty" reform:"mtime"`
}

type Responder struct {
	code     int
	response interface{}
	headers  http.Header
}

func (r *Responder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	for k, v := range r.headers {
		for _, val := range v {
			rw.Header().Add(k, val)
		}
	}

	rw.WriteHeader(r.code)

	if err := producer.Produce(rw, r.response); err != nil {
		panic(err)
	}
}

func ListAccountsHandlerFunc(params account.ListAccountsParams) middleware.Responder {
	result, err := db.Querier.SelectAllFrom(AccountTable, "")
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	return &Responder{http.StatusOK, result, make(http.Header)}
}

func AddAccountHandlerFunc(params account.AddAccountParams) middleware.Responder {
	rec := Account{Name: params.Name, Amount: params.Amount}
	err := db.Querier.Insert(&rec)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	return &Responder{http.StatusOK, &rec, make(http.Header)}
}

func DeleteAccountHandlerFunc(params account.DeleteAccountParams) middleware.Responder {
	rec := Account{ID: params.ID}
	err := db.Querier.Delete(&rec)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	return &Responder{http.StatusNoContent, nil, make(http.Header)}
}

func DepositToAccountHandlerFunc(params account.DepositToAccountParams) (res middleware.Responder) {
	commit := false

	tx, err := db.Begin()
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	defer func() {
		if commit {
			err = tx.Commit()
		} else {
			err = tx.Rollback()
		}
		if err != nil {
			res = &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
		}
	}()

	_, err = tx.Querier.Exec(
		fmt.Sprintf(
			"UPDATE %s SET amount = amount + ? WHERE id = ?",
			AccountTable.Name(),
		),
		params.Amount,
		params.ID,
	)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	rec, err := tx.Querier.FindOneFrom(AccountTable, "id", params.ID)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	commit = true
	return &Responder{http.StatusOK, &rec, make(http.Header)}
}

func FindAccountByNameHandlerFunc(params account.FindAccountByNameParams) middleware.Responder {
	rec, err := db.Querier.FindOneFrom(AccountTable, "name", params.Name)
	if err != nil {
		if err == reform.ErrNoRows {
			return &Responder{http.StatusNotFound, err.Error(), make(http.Header)}
		}
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	return &Responder{http.StatusOK, &rec, make(http.Header)}
}

func GetAccountByIDHandlerFunc(params account.GetAccountByIDParams) middleware.Responder {
	rec, err := db.Querier.FindOneFrom(AccountTable, "id", params.ID)
	if err != nil {
		if err == reform.ErrNoRows {
			return &Responder{http.StatusNotFound, err.Error(), make(http.Header)}
		}
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	return &Responder{http.StatusOK, &rec, make(http.Header)}
}

func TransferBetweenAccountsHandlerFunc(params account.TransferBetweenAccountsParams) (res middleware.Responder) {
	commit := false

	tx, err := db.Begin()
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	defer func() {
		if commit {
			err = tx.Commit()
		} else {
			err = tx.Rollback()
		}
		if err != nil {
			res = &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
		}
	}()

	_, err = tx.Querier.Exec(
		fmt.Sprintf(
			"UPDATE %s SET amount = amount - ? WHERE id = ?",
			AccountTable.Name(),
		),
		params.Amount,
		params.ID,
	)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	_, err = tx.Querier.Exec(
		fmt.Sprintf(
			"UPDATE %s SET amount = amount + ? WHERE id = ?",
			AccountTable.Name(),
		),
		params.Amount,
		params.Target,
	)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	recs := make([]reform.Struct, 2)
	recs[0], err = tx.Querier.FindOneFrom(AccountTable, "id", params.ID)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	recs[1], err = tx.Querier.FindOneFrom(AccountTable, "id", params.Target)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	commit = true
	return &Responder{http.StatusOK, recs, make(http.Header)}
}

func WithdrawAccountHandlerFunc(params account.WithdrawAccountParams) (res middleware.Responder) {
	commit := false

	tx, err := db.Begin()
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}
	defer func() {
		if commit {
			err = tx.Commit()
		} else {
			err = tx.Rollback()
		}
		if err != nil {
			res = &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
		}
	}()

	_, err = tx.Querier.Exec(
		fmt.Sprintf(
			"UPDATE %s SET amount = amount - ? WHERE id = ?",
			AccountTable.Name(),
		),
		params.Amount,
		params.ID,
	)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	rec, err := tx.Querier.FindOneFrom(AccountTable, "id", params.ID)
	if err != nil {
		return &Responder{http.StatusInternalServerError, err.Error(), make(http.Header)}
	}

	commit = true
	return &Responder{http.StatusOK, &rec, make(http.Header)}
}
