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
		seelog.Critical("日志配置有问题:", err)
		return
	}
	seelog.ReplaceLogger(logger)

	defer seelog.Flush()

	beego.SetStaticPath("/static","static")
	beego.SetStaticPath("/ueditor","static/ueditor")

}

func main() {

	beego.Run()
}
