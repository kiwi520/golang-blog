package routers

import (
	"bloggo/controllers"
	"bloggo/controllers/index"
	"bloggo/controllers/admin"
	"github.com/astaxie/beego"
	"bloggo/controllers/apis"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))


	beego.Router("/user", &controllers.UserController{})
	beego.Router("/index", &index.IndexController{})

	//百度编辑器ueditor
	beego.Router("/ueidtor", &controllers.UeditorController{},"*:Ueditor")

	admin :=
		beego.NewNamespace("/admin",
			beego.NSRouter("/", &admin.IndexController{}),
			beego.NSRouter("/success", &admin.SuccessController{}),
			beego.NSNamespace("/article",
				beego.NSRouter("/list", &admin.ArticleController{}, "get:ListArticle"),
				beego.NSRouter("/add", &admin.ArticleController{}, "get:AddArticle;post:AddArticle"),
				beego.NSRouter("/detail/:id", &admin.ArticleController{}, "get:GetDetailArticle"),
				beego.NSRouter("/update/?:id", &admin.ArticleController{}, "get:UpdateArticle;post:UpdateArticle"),
			),
		)
	//注册 后台路由
	beego.AddNamespace(admin)

	api :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/article",
				beego.NSRouter("/list", &apis.ArticleController{}, "get:ListArticle"),
				beego.NSRouter("/detail/:id", &apis.ArticleController{}, "get:GetDetailArticle"),
			),
		)
	//注册 API路由
	beego.AddNamespace(api)
}
