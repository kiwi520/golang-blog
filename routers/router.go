package routers

import (
	"bloggo/controllers"
	"bloggo/controllers/index"
	"bloggo/controllers/admin"
	"github.com/astaxie/beego"
	//"github.com/derekparker/delve/service/api"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/index", &index.IndexController{})
	beego.Router("/admin", &admin.IndexController{})


	ns :=
		beego.NewNamespace("/go",
			beego.NSRouter("/ss", &controllers.MainController{}),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
