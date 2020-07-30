package mongodb

import (
	"fmt"

	dbcfg "../config"
	mgo "gopkg.in/mgo.v2"
)

var config = dbcfg.MongoConfig{}
var (
	Session *mgo.Session
	DB      *mgo.Database
)

func init() {
	CreateSession()
}

// Parse the configuration file 'config-mongo.toml', and establish a connection to DB
func CreateSession() {
	config.ReadConfig()

	var err error
	Session, err = mgo.Dial(config.Server)
	if err != nil {
		panic(err)
	}

	DB = Session.DB(config.Database)
	DB.Login(config.Username, config.Password)

	fmt.Println("SESSION ACTIVE")
}
