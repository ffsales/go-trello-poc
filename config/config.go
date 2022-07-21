package config

import (
	"fmt"

	"github.com/magiconair/properties"
)

type configuration struct {
	DbConfig DB
}

type DB struct {
	Host     string `properties:"host",default="local"`
	Port     string `properties:"port",default="3306"`
	User     string `properties:"user",default="root"`
	Pass     string `properties:"pass",default="123456"`
	Database string `properties:"database_name",default="go-trello"`
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
		User:     pf.MustGet("port"),
		Pass:     pf.MustGet("pass"),
		Database: pf.MustGet("database_name"),
	}
}

func GetDBConfig() DB {
	return conf.DbConfig
}
