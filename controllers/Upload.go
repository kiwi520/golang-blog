package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"fmt"
)

/**
 * 文件描述
 * Create on 17-12-12 上午11:46
 * Create by hu
 */

type UploadController struct {
	beego.Controller
}

func (c *UploadController) UploadImg() {

	path :="static/upload/markdown/images/"

	err := os.MkdirAll(path, 0777)  //给目录符权限
	if err != nil {
		beego.Error(err)
	}

	f, h, _ := c.GetFile("txt_file")                  //获取上传的文件
	fmt.Println(h.Size)
	//fmt.Println(f.Read)
	fullpath := path + h.Filename
	f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	c.SaveToFile("txt_file", fullpath)                    //存文件
	c.Data["json"] =map[string]interface{}{"state": "SUCCESS", "url": fullpath, "title": h.Filename, "original": h.Filename}
	c.ServeJSON()
}






//检查目录是否存在
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Print(filename + " not exist")
		exist = false
	}
	return exist
}