package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"bloggo/models"
	"log"
	"github.com/astaxie/beego/orm"
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

func(this *LoginController) GetLogin(){

	this.TplName = "admin/login/login.html"
}

func(this *LoginController) PostLogin(){
		u := models.User{}
		if err := this.ParseForm(&u); err != nil {
			//handle error
		}

		valid := validation.Validation{}
		b, err := valid.Valid(&u)
		m := make(map[string]string)
		if err != nil {
			// handle error
		}
		if !b {
			for _, err := range valid.Errors {
				m[err.Key]=err.Message
				log.Println(err.Key, err.Message)
			}
		}
			o := orm.NewOrm()
			var user models.User
			ps := o.QueryTable("user")
			errss :=ps.Filter("email",u.Email).One(&user)
			if errss ==nil {
				sess := this.StartSession()
				sess.Set("uid",user.Id)
				sess.Set("email",user.Email)
				this.TplName = "admin/welcome.html"

				if len(m) <=0 {
					this.Data["ok"] = "success"
				}else {
					this.Data["ok"] =""
				}
				this.Data["err"] = m
				this.Ctx.Redirect(302, "/admin/")
			}
}


func (this *BaseController) Prepare() {
	sess_uid :=this.GetSession("uid")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
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


