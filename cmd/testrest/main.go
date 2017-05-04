package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/onokonem/go-testrest/config"

	"github.com/onokonem/go-testrest/proto/restapi"
	"github.com/onokonem/go-testrest/proto/restapi/operations"
)

func main() {
	config.ParseFlags()
	err := config.ParseConf(*config.Flags.Config)
	if err != nil {
		panic(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTestrestAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = config.Conf.Host
	server.Port = config.Conf.Port
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
