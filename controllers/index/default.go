package index

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
	this.Data["HtmlHead"]="header"
	this.Data["LayoutContent"]="LayoutContent"
	this.Data["SideBar"]="SideBar"
	this.TplName = "layout/layout.html"
	//this.Layout = "layout/layout.html"
}