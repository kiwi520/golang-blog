package admin

/**
 * 文件描述
 * Create on 11/28/17 2:17 PM
 * Create by hu
 */
import (
	"bloggo/models"
	"github.com/astaxie/beego/orm"
	"github.com/cihub/seelog"
)


type SuccessController struct{
	BaseController
}


type IndexController struct {
	BaseController
}
type ArticleController struct {
	BaseController
}


func (this *SuccessController) Get() {
	this.TplName="admin/success.html"
}


/**
* 首页
 */
func (this *IndexController) Get() {
	//this.LayoutSections["HtmlHead"]="header"
	//this.LayoutSections["LayoutContent"]="admin/welcome.html"
	//this.LayoutSections["SideBar"]="SideBar"
	//this.Data["Context"] = "value"
	//this.TplName = "layout/layout.html"
	this.TplName = "admin/welcome.html"
	this.Layout = "admin/layout.html"

}


/**
*添加文章
 */
func (this *ArticleController) AddArticle() {
	this.TplName = "admin/article/add.html"
	this.Layout = "admin/layout.html"
	if this.Ctx.Request.Method == "POST" {
		art := new(models.Article)
		o := orm.NewOrm()
		if err := this.ParseForm(art); err != nil {
			seelog.Error("文章添加失败")
		}else{
			_, err := o.Insert(art)
			if err == nil {
				seelog.Info("文章添加成功")
				this.Success("提交成功","/admin/success",1)
			}else{
				this.Error("提交失败","",2)
			}
		}
	}
}
