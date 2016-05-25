package mysqlCore

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"Go-Rubix/core/common"
)


func Init(){
	db, err := sql.Open("mysql", "root:mypassword@/rubix")
	common.CheckErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect:gorp.MySQLDialect{"InnoDB", "UTF8"}}
	return dbmap
}

