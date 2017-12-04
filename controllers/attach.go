package controllers

import (
	"github.com/astaxie/beego"
	"net/url"
	"os"
	"io"
)

/**
 * 文件描述
 * Create on 17-12-4 上午10:26
 * Create by hu
 */

type AttachController struct {
	beego.Controller
}

func (c *AttachController) Get() {
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:])
	_, err = url.QueryUnescape(c.Ctx.Request.RequestURI)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(c.Ctx.ResponseWriter, f)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
}