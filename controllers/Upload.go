package controllers

import (
	"github.com/astaxie/beego"
	"github.com/muesli/smartcrop"
	"os"
	"fmt"
	"image"
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
	//fmt.Println(h.Size)

	fullpath := path + h.Filename
	f.Close()                                          //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	c.SaveToFile("txt_file", fullpath)
	//存文件
	hostname := c.Ctx.Input.Site()
	partname := beego.AppConfig.String("httpport")
	c.Data["json"] =map[string]interface{}{"state": "SUCCESS","hostname":hostname,"partname":partname, "url": fullpath, "title": h.Filename, "original": h.Filename}
	c.ServeJSON()
}


func cropimg (urlimg string){
	f, _ := os.Open(urlimg)
	img, _, _ := image.Decode(f)

	analyzer := smartcrop.NewAnalyzer()
	topCrop, _ := analyzer.FindBestCrop(img, 250, 250)

	// The crop will have the requested aspect ratio, but you need to copy/scale it yourself
	fmt.Printf("Top crop: %+v\n", topCrop)

	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	croppedimg := img.(SubImager).SubImage(topCrop)

	fmt.Println(croppedimg)
}