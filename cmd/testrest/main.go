package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/mysql"

	"github.com/go-openapi/loads"
	"github.com/onokonem/go-testrest/config"

	"github.com/onokonem/go-testrest/proto/restapi"
	"github.com/onokonem/go-testrest/proto/restapi/operations"
	"github.com/onokonem/go-testrest/proto/restapi/operations/account"
)

var db *reform.DB

func main() {
	config.ParseFlags()
	err := config.ParseConf(*config.Flags.Config)
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(config.Conf.DB.Driver, config.Conf.DB.DSN)
	if err != nil {
		panic(err)
	}

	db = reform.NewDB(conn, mysql.Dialect, reform.NewPrintfLogger(log.New(os.Stderr, "SQL: ", log.Flags()).Printf))

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTestrestAPI(swaggerSpec)

	// Register my implementation
	api.AccountAddAccountHandler = account.AddAccountHandlerFunc(AddAccountHandlerFunc)
	api.AccountDeleteAccountHandler = account.DeleteAccountHandlerFunc(DeleteAccountHandlerFunc)
	api.AccountDepositToAccountHandler = account.DepositToAccountHandlerFunc(DepositToAccountHandlerFunc)
	api.AccountFindAccountByNameHandler = account.FindAccountByNameHandlerFunc(FindAccountByNameHandlerFunc)
	api.AccountGetAccountByIDHandler = account.GetAccountByIDHandlerFunc(GetAccountByIDHandlerFunc)
	api.AccountListAccountsHandler = account.ListAccountsHandlerFunc(ListAccountsHandlerFunc)
	api.AccountTransferBetweenAccountsHandler = account.TransferBetweenAccountsHandlerFunc(TransferBetweenAccountsHandlerFunc)
	api.AccountWithdrawAccountHandler = account.WithdrawAccountHandlerFunc(WithdrawAccountHandlerFunc)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = config.Conf.Host
	server.Port = config.Conf.Port

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
