// testrest configuration, by cmd flags and YAML file
package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// DB driver definition struct
// Driver: DB driver name as it will be bassed to sql.Open(driverName, dataSourceName string)
// DSN: DB DSN (Data Source Name) as it will be bassed to sql.Open(driverName, dataSourceName string)
type StructConfigDB struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"db"`
}

// Struct representation for a YAML config file
// Listen: local [addr]:port to listen on
// DB: database params (see StructConfigDB)
type StructConfig struct {
	Host string         `yaml:"host"`
	Port int            `yaml:"port"`
	DB   StructConfigDB `yaml:"db"`
}

// Config struct ready to be utilized by yaml.Unmarshal(in []byte, out interface{})
// with all the defaults set
var Conf = StructConfig{
	Host: "0.0.0.0",
	Port: 8080,
	DB: StructConfigDB{
		Driver: "mysql",
		DSN:    "root@/testrest",
	},
}

// Simple YAML unmarshaler wrapper
// Note: file absense is not a error at all
func ParseConf(file string) error {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Failed to read file %q: %v", file, err)
	}

	err = yaml.Unmarshal(content, &Conf)
	if err != nil {
		return fmt.Errorf("Failed to parse file %q: %v", file, err)
	}

	return nil
}
