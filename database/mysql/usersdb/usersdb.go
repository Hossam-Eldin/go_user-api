package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	//for mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Client Db
var (
	Client   *sql.DB
	username = "root"
	password = ""
	host     = "localhost:3306"
	schema   = "go_users"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("========== database is configerued ==========")

}
