package ormscripts

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type config struct {
	user   string
	pass   string
	addr   string
	port   string
	dbname string
}

func db_init(db_name string) *gorm.DB {
	conf := &config{
		user:   "root",
		pass:   "dong20020324",
		addr:   "127.0.0.1",
		port:   "3306",
		dbname: db_name,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.user, conf.pass, conf.addr, conf.port, conf.dbname)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
