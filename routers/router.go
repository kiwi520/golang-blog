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
	admin :=
		beego.NewNamespace("/admin",
			beego.NSRouter("/", &admin.IndexController{}),
			beego.NSNamespace("/article",
				beego.NSRouter("/add", &admin.ArticleController{}, "get:AddArticle;post:AddArticle"),
			),
		)
	//注册 namespace
	beego.AddNamespace(admin)
}
