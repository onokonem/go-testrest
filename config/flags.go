// testrest configuration, by cmd flags and YAML file
package config

import (
	"flag"
)

type structFlags struct {
	Config *string
}

// command line flags ready to be utilized by flag.Parse()
var Flags = structFlags{
	Config: flag.String("config", "./testrest.yml", "Path to a config file"),
}

// Dummy flag.Parse() wrapper
func ParseFlags() {
	flag.Parse()
}
