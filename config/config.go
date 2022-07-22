package config

import (
	"fmt"

	"github.com/magiconair/properties"
)

type configuration struct {
	DbConfig DB
}

type DB struct {
	Host     string `properties:"host"`
	Port     string `properties:"port"`
	User     string `properties:"user"`
	Pass     string `properties:"pass"`
	Database string `properties:"database_name"`
}

var conf *configuration

func init() {
	properties.ErrorHandler = func(err error) {
		fmt.Println(err)
		fmt.Println("Loadding default values")
	}
	pf := properties.MustLoadFile("${HOME}/workspace/go-trello-poc/static/config.properties", properties.UTF8)

	conf = new(configuration)

	conf.DbConfig = DB{
		Host:     pf.MustGet("host"),
		Port:     pf.MustGet("port"),
		User:     pf.MustGet("user"),
		Pass:     pf.MustGet("pass"),
		Database: pf.MustGet("database_name"),
	}
}

func GetDBConfig() DB {
	return conf.DbConfig
}
