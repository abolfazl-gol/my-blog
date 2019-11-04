package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type Engine interface {
	Table(tableNameOrBean interface{}) *xorm.Session
	Select(str string) *xorm.Session
	Count(...interface{}) (int64, error)
	Decr(column string, arg ...interface{}) *xorm.Session
	Delete(interface{}) (int64, error)
	Exec(...interface{}) (sql.Result, error)
	Exist(bean ...interface{}) (bool, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	ID(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Incr(column string, arg ...interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Join(joinOperator string, tablename interface{}, condition string, args ...interface{}) *xorm.Session
	SQL(interface{}, ...interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
	Asc(colNames ...string) *xorm.Session
}

var (
	engine *xorm.Engine
)

func InitDB() {
	var err error
	engine, err = xorm.NewEngine("mysql", os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
	}

	engine.SetColumnMapper(core.GonicMapper{})
}

func CloseDB() {
	engine.Close()
}
