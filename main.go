package main

import (
	_ "bloggo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"bloggo/models"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.13.106:3306)/go?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.Article))

	// create table
	//orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
