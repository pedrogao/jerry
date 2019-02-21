package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
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
	//log.Println(config)
	engine, err := xorm.NewEngine("mysql", config)
	if err != nil {
		log.Fatalf("%s, Database connection failed. Database name: %s", err, name)
	}
	err = engine.Ping()
	if err != nil {
		log.Fatalf("%s, Database is killed. Database name: %s", err, name)
	}
	setupDB(engine)
	return engine
}

func setupDB(db *xorm.Engine) {
	db.SetLogLevel(core.LOG_DEBUG)
	//db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	//db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
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
