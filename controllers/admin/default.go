package admin

/**
 * 文件描述
 * Create on 11/28/17 2:17 PM
 * Create by hu
 */
import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}


func (this *IndexController) Get() {
	//this.LayoutSections["HtmlHead"]="header"
	//this.LayoutSections["LayoutContent"]="admin/content.html"
	//this.LayoutSections["SideBar"]="SideBar"
	//this.Data["Context"] = "value"
	//this.TplName = "layout/layout.tpl"
	this.TplName = "admin/header.html"
	this.Layout = "layout/layout.tpl"

}