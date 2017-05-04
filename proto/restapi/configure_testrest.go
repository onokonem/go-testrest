package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/onokonem/go-testrest/proto/restapi/operations"
	"github.com/onokonem/go-testrest/proto/restapi/operations/account"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yml

func configureFlags(api *operations.TestrestAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TestrestAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AccountAddAccountHandler = account.AddAccountHandlerFunc(func(params account.AddAccountParams) middleware.Responder {
		return middleware.NotImplemented("operation account.AddAccount has not yet been implemented")
	})
	api.AccountDeleteAccountHandler = account.DeleteAccountHandlerFunc(func(params account.DeleteAccountParams) middleware.Responder {
		return middleware.NotImplemented("operation account.DeleteAccount has not yet been implemented")
	})
	api.AccountDepositToAccountHandler = account.DepositToAccountHandlerFunc(func(params account.DepositToAccountParams) middleware.Responder {
		return middleware.NotImplemented("operation account.DepositToAccount has not yet been implemented")
	})
	api.AccountFindAccountByNameHandler = account.FindAccountByNameHandlerFunc(func(params account.FindAccountByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation account.FindAccountByName has not yet been implemented")
	})
	api.AccountGetAccountByIDHandler = account.GetAccountByIDHandlerFunc(func(params account.GetAccountByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation account.GetAccountByID has not yet been implemented")
	})
	api.AccountListAccountsHandler = account.ListAccountsHandlerFunc(func(params account.ListAccountsParams) middleware.Responder {
		return middleware.NotImplemented("operation account.ListAccounts has not yet been implemented")
	})
	api.AccountTransferBetweenAccountsHandler = account.TransferBetweenAccountsHandlerFunc(func(params account.TransferBetweenAccountsParams) middleware.Responder {
		return middleware.NotImplemented("operation account.TransferBetweenAccounts has not yet been implemented")
	})
	api.AccountWithdrawAccountHandler = account.WithdrawAccountHandlerFunc(func(params account.WithdrawAccountParams) middleware.Responder {
		return middleware.NotImplemented("operation account.WithdrawAccount has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
