package admin

/**
 * 文件描述
 * Create on 11/28/17 2:17 PM
 * Create by hu
 */
import (
	"github.com/astaxie/beego"
	//"bloggo/models"
	"fmt"
	"bloggo/models"
	"github.com/astaxie/beego/orm"
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
		art := article{}
		Art := new(models.Article)
		o := orm.NewOrm()
		if err := this.ParseForm(&art); err != nil {
			//fmt.Println(art)
		}else{
			Art.Title =art.Title
			Art.Content =art.Content
			id, err := o.Insert(&Art)
			if err == nil {
				fmt.Println(id)
			}
		}
	}
}