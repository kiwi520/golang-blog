package main

import (
	_ "bloggo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"bloggo/models"
	"github.com/cihub/seelog"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.13.106:3306)/homestead?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.Article))

	// create table
	//orm.RunSyncdb("default", false, true)

	logger, err := seelog.LoggerFromConfigAsFile("conf/seelog-dev-main.xml")

	if err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}
	seelog.ReplaceLogger(logger)

	defer seelog.Flush()

}

func main() {
	beego.Run()
}
