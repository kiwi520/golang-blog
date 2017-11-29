package admin

/**
 * 文件描述
 * Create on 11/28/17 2:17 PM
 * Create by hu
 */
import (
	"github.com/astaxie/beego"
	//"fmt"
	//"bloggo/models"
)

type IndexController struct {
	beego.Controller
}
type ArticleController struct {
	beego.Controller
}

type article struct {
	Id    int         `form:"-"`
	Title  string		 `form:"title"`
	Content   string         `form:"content"`
}


func (this *IndexController) Get() {
	//this.LayoutSections["HtmlHead"]="header"
	//this.LayoutSections["LayoutContent"]="admin/welcome.html"
	//this.LayoutSections["SideBar"]="SideBar"
	//this.Data["Context"] = "value"
	//this.TplName = "layout/layout.html"
	this.TplName = "admin/welcome.html"
	this.Layout = "admin/layout.html"

}

func (this *ArticleController) AddArticle() {
	this.TplName = "admin/article/add.html"
	this.Layout = "admin/layout.html"
	if this.Ctx.Request.Method == "POST" {
		//title := this.GetString("title")
		//content := this.GetString("content")
		////art := article{}
		//fmt.Print(title)
		//fmt.Print(content)
		//this.Data["User"] = user
		//this.Redirect("/", 302)
		//art := models.Article{}
		//if err := this.ParseForm(&art); err != nil {
		//	fmt.Println(art)
		//}else{
		//	fmt.Println(art)
		//}
	}
}
//func (this *ArticleController) GetAddArticle() {
//	this.TplName = "admin/article/add.html"
//	this.Layout = "admin/layout.html"
//}
//
//func (this *ArticleController) PostAddArticle() {
//	art := article{}
//	if err := this.ParseForm(&art); err != nil {
//		fmt.Print(art)
//	}
//	//jsoninfo := this.GetString("jsoninfo")
//	//if jsoninfo == "" {
//	//	this.Ctx.WriteString("jsoninfo is empty")
//	//	return
//	//}
//}