package entities

import (
	//"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//var mydb *sql.DB
var engine *xorm.Engine

func init() {
	//create engine
	dbEngine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	err = dbEngine.Sync2(new(UserInfo))
	checkErr(err)
	engine = dbEngine
}

// DaoSource Data Access Object Source
//type DaoSource struct {
// if DB, each statement execute sql with random conn.
// if Tx, all statements use the same conn as the Tx's connection
//	SQLExecer
//}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
