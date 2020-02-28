package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MySqlDB *sql.DB //MySQL

func init() {
	var err error
	connUrl := "root:123456@tcp(localhost:3306)/go_demo"
	MySqlDB, err = sql.Open("mysql", connUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	MySqlDB.SetMaxIdleConns(20)
	MySqlDB.SetMaxOpenConns(20)

	if err := MySqlDB.Ping(); err != nil {
		log.Fatalln(err.Error())
	}

}
