package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DB *xorm.Engine

func openDB(username, password, addr, name string) *xorm.Engine {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	engine, err := xorm.NewEngine("mysql", config)
	if err != nil {
		log.Fatalf("%s, Database connection failed. Database name: %s \n", err, name)
	}
	err = engine.Ping()
	if err != nil {
		log.Fatalf("%s, Database is killed. Database name: %s \n", err, name)
	}
	setupDB(engine)
	return engine
}

func setupDB(db *xorm.Engine) {
	db.SetLogLevel(core.LOG_DEBUG)
}

func InitLocal() *xorm.Engine {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func Init() {
	DB = InitLocal()
}

func Close() {
	DB.Close()
}
