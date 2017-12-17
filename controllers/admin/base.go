package admin

import (
	"github.com/astaxie/beego"
	//"net/http"
)

/**
 * 文件描述
 * Create on 11/30/17 1:32 PM
 * Create by hu
 */


/**
* login
 */
type LoginController struct{
	beego.Controller
}
/**
*利用结构体嵌套实现继承 Beego controller里的方法
 */
type BaseController struct{
	beego.Controller
}

func(this *LoginController) Login(){
	this.TplName = "admin/login/login.html"
	//this.Layout = "admin/layout.html"
}


//func (this *BaseController) Prepare(w http.ResponseWriter, r *http.Request) {
//	sess := globalSessions.SessionStart(w, r)
//	defer sess.SessionRelease()
//	sess_uid := sess.Get("userid")
//	sess_username := sess.Get("username")
//	if sess_uid == nil {
//		this.Ctx.Redirect(302, "/admin/login")
//		return
//	}
//	this.Data["Username"] = sess_username
//}



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


