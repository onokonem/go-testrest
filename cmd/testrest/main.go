package main

import (
	"github.com/onokonem/go-testrest/config"
)

func main() {
	config.ParseFlags()
	err := config.ParseConf(*config.Flags.Config)
	if err != nil {
		panic(err)
	}
}
