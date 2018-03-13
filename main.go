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
	dbaddr := beego.AppConfig.String("dbaddr")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbname := beego.AppConfig.String("dbname")
	dbpasswd := beego.AppConfig.String("dbpasswd")
	orm.RegisterDataBase("default", "mysql", dbuser+":"+dbpasswd+"@tcp("+dbaddr+":"+dbport+")/"+dbname+"?charset=utf8&loc=Asia%2FShanghai", 30)

	// register model
	orm.RegisterModel(new(models.Article),new(models.User))

	// create table
	//orm.RunSyncdb("default", false, true)
	//session start
	beego.BConfig.WebConfig.Session.SessionOn = true

	logger, err := seelog.LoggerFromConfigAsFile("conf/seelog-dev-main.xml")

	if err != nil {
		seelog.Critical("日志配置有问题:", err)
		return
	}
	seelog.ReplaceLogger(logger)

	defer seelog.Flush()

	//beego.SetStaticPath("/static","static")
	beego.SetStaticPath("/ueditor","static/ueditor")
	beego.SetStaticPath("/markdown","static/markdown")
	beego.SetStaticPath("/editormd","static/editormd")
	beego.SetStaticPath("/fileinput","static/bootstrap-fileinput")
	beego.SetStaticPath("/datetimepicker","static/datetimepicker")

}

func main() {
	beego.Run()
}