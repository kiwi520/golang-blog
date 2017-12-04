package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"github.com/cihub/seelog"
	"io/ioutil"
	"regexp"
	"encoding/json"
)

/**
 * 文件描述
 * Create on 17-12-4 上午8:43
 * Create by hu
 */

type UeditorController struct {
	beego.Controller
}



func (c *UeditorController) Ueditor() {
	op := c.GetString("action")
	var fileName string
	if op =="uploadimage"{
		fileName = "images"
	} else if op =="uploadfile"{
		fileName = "files"
	}
    switch op {
	case "config":
		file,err := os.Open("conf/config.json")
		if err !=nil{
			seelog.Warn("获取ueditorq配置文件出错:"+err.Error())
		}

		defer file.Close()

		fd, err := ioutil.ReadAll(file)
		src := string(fd)
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/")
		src = re.ReplaceAllString(src, "")
		tt := []byte(src)
		var r interface{}
		json.Unmarshal(tt, &r) //这个byte要解码/
		c.Data["json"] = r
		c.ServeJSON()
	case "uploadimage", "uploadfile", "uploadvideo":
		number := c.GetString("number")
		//name := c.GetString("name")
		err := os.MkdirAll("static/upload/"+ fileName +"/", 0777)  //给目录符权限
		if err != nil {
			beego.Error(err)
		}
		//保存上传的图片
		//获取上传的文件，直接可以获取表单名称对应的文件名，不用另外提取
		_, h, err := c.GetFile("upfile")
		if err != nil {
			beego.Error(err)
		}
		pathPath := "static/upload/"+ fileName  + number+ "/" + h.Filename
		err = c.SaveToFile("upfile", pathPath)
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": "/static/upload/"+ fileName  + number +"/" + h.Filename, "title": h.Filename, "original": h.Filename}
		c.ServeJSON()
	}
}