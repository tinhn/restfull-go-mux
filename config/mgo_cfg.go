package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type MongoConfig struct {
	Server   string
	Database string
	Username string
	Password string
}

// Read and parse the configuration file
func (c *MongoConfig) ReadConfig() {
	if _, err := toml.DecodeFile("mongo-cfg.toml", &c); err != nil {
		log.Fatal(err)
	}
}
