package routers

import (
	"bloggo/controllers"
	"bloggo/controllers/index"
	"bloggo/controllers/admin"
	"github.com/astaxie/beego"
	"bloggo/controllers/apis"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/index", &index.IndexController{})
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
				//beego.NSRouter("/update/", &admin.ArticleController{}, "post:UpdateArticle"),
			),
		)
	//注册 namespace
	beego.AddNamespace(admin)




	api :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/article",
				beego.NSRouter("/list", &apis.ArticleController{}, "get:ListArticle"),
				beego.NSRouter("/detail/:id", &apis.ArticleController{}, "get:GetDetailArticle"),
				//beego.NSRouter("/update/", &admin.ArticleController{}, "post:UpdateArticle"),
			),
		)
	//注册 namespace
	beego.AddNamespace(api)
}
