package db

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/jackyyf/gook/utils/log"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var db *sql.DB

func quoteSingle(str string) string {
	return strings.Replace(strings.Replace(str, "\\", "\\\\", -1), "'", "\\'", -1)
}

func init() {
	log.Debug("Call: db.init")
	conf, err := config.NewConfig("ini", "conf/db.ini")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open database config failed: %s", err)
		panic(err)
	}
	host := quoteSingle(conf.DefaultString("host", "localhost"))
	port := conf.DefaultInt("port", 5432)
	user := quoteSingle(conf.DefaultString("user", "root"))
	password := quoteSingle(conf.DefaultString("passwd", ""))
	dbname := quoteSingle(conf.DefaultString("db", "gook"))
	timeout := conf.DefaultInt("timeout", 0)
	conn_str := fmt.Sprintf("host='%s' port=%d user='%s' password='%s' dbname='%s' connect_timeout=%d", host, port, user, password, dbname, timeout)
	log.Info("Conn string: %s", conn_str)
	db, err = sql.Open("postgres", conn_str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %s", err)
		panic(err)
	}
	log.Debug("Done: db.init")
}

func Transaction() (*sql.Tx, error) {
	log.Debug("db: transcation")
	return db.Begin()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Debug("db: exec %s %v", query, args)
	return db.Exec(query, args...)
}

func Prepare(query string) (*sql.Stmt, error) {
	log.Debug("db: prepare %s", query)
	return db.Prepare(query)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Debug("db: query %s %v", query, args)
	return db.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	log.Debug("db: queryrow %s %v", query, args)
	return db.QueryRow(query, args...)
}
