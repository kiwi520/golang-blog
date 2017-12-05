package apis

import (
	"bloggo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
 * 文件描述 : API接口
 * Create on 17-12-5 下午2:36
 * Create by hu
 */

type ArticleController struct{
	beego.Controller
}


func (c *ArticleController) ListArticle (){
	o := orm.NewOrm()
	var info []models.Article
	nums,err :=o.QueryTable("article").All(&info)
	if err ==nil && nums >0 {
		c.Data["json"] = map[string]interface{}{"status_code":200,"status_msg": "SUCCESS", "data":info}
		c.ServeJSON()
	}else{
		c.Data["json"] = map[string]interface{}{"status_code":404,"status_msg": "no found data", "data":""}
		c.ServeJSON()
	}
}

func (c *ArticleController) GetDetailArticle(){
	c.Ctx.WriteString("apis")
}