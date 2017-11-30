package admin

/**
 * 文件描述
 * Create on 11/28/17 2:17 PM
 * Create by hu
 */
import (
	"github.com/astaxie/beego"
	"bloggo/models"
	"github.com/astaxie/beego/orm"
	"github.com/cihub/seelog"
)

type BaseController struct{
	beego.Controller
}


type SuccessController struct{
	BaseController
}


type IndexController struct {
	BaseController
}
type ArticleController struct {
	BaseController
}



type article struct {
	Id    int         `form:"-"`
	Title  string		 `form:"title"`
	Content   string         `form:"content"`
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





/*
* 成功跳转
 */
func (this *BaseController) Success(msg string, url string, wait int) {
	data := make(map[string]interface{})
	data["type"] = true
	data["title"] = "提示信息"
	data["msg"] = msg
	data["wait"] = wait
	if url == "-1" {
		url = this.Ctx.Request.Referer()
	} else if url == "-2" {
		url = this.Ctx.Request.Referer()
	}
	data["url"] = url
	this.Data["mess"] = data
	this.TplName = "message.html"

}

/*
* 失败跳转
 */
func (this *BaseController) Error(msg string, url string, wait int) {
	data := make(map[string]interface{})
	data["type"] = false
	data["title"] = "错误提示"
	data["msg"] = msg
	data["wait"] = wait
	if url == "-1" {
		url = this.Ctx.Request.Referer()
	} else if url == "-2" {
		url = this.Ctx.Request.Referer()
	}

	data["url"] = url
	this.Data["mess"] = data
	this.TplName = "message.html"

}

/*
* Ajax返回
*
 */
func (this *BaseController) AjaxReturn(status int, msg string, data interface{}) {
	json := make(map[string]interface{})
	json["status"] = status
	json["msg"] = msg
	json["data"] = data
	this.Data["json"] = json
	this.ServeJSON()
	return
}

