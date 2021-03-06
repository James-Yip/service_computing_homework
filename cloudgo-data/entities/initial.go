package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// var mydb *sql.DB

func init() {
	// initialize xorm engine
	Engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	Engine.SetMapper(core.SameMapper{})
	err = Engine.Sync2(new(UserInfo))
	checkErr(err)
	engine = Engine
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
