// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/muzzlefork/first-api/restapi/operations"
)

//go:generate swagger generate server --target ../../first-api --name TrfMicroserviceGolangFrb --spec ../swagger.yml

func configureFlags(api *operations.TrfMicroserviceGolangFrbAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TrfMicroserviceGolangFrbAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.Oauth2Auth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (oauth2) has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.GetConfigHandler == nil {
		api.GetConfigHandler = operations.GetConfigHandlerFunc(func(params operations.GetConfigParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetConfig has not yet been implemented")
		})
	}
	if api.GetHealthHandler == nil {
		api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation .GetHealth has not yet been implemented")
		})
	}

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
func configureServer(s *http.Server, scheme, addr string) {
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
