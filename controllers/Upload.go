package controllers

import (
	"github.com/astaxie/beego"
	"code.google.com/p/graphics-go/graphics"
	"os"
	"fmt"
	"image"
	"github.com/cihub/seelog"
	"image/png"
	//"strings"
	"strings"
)

/**
 * 文件描述
 * Create on 17-12-12 上午11:46
 * Create by hu
 */

type UploadController struct {
	beego.Controller
}

/**
* 上传图片
 */
func (c *UploadController) UploadImg() {

	path :="static/upload/markdown/images/"

	err := os.MkdirAll(path, 0777)  //给目录符权限
	if err != nil {
		seelog.Warn("给上传图片目录设置权限失败")
	}

	f, h, _ := c.GetFile("txt_file")                  //获取上传的文件

	fullpath := path + h.Filename
	f.Close()                                          //关闭上传的文件
	c.SaveToFile("txt_file", fullpath)
	//存文件
	thumbnail :=cropimg(path,h.Filename,5)
	hostname := c.Ctx.Input.Site()
	partname := beego.AppConfig.String("httpport")
	c.Data["json"] =map[string]interface{}{"state": "SUCCESS","hostname":hostname,"partname":partname, "url": fullpath, "thumb":thumbnail,"title": h.Filename, "original": h.Filename}
	c.ServeJSON()
}

/**
*图片缩略
 */
func cropimg (path,filename string,num int) (aimspath string){
	//原图片路径
	fullpath := path + filename

	//aimspath = path +"cron/"
	//目标路径
	lens := strings.LastIndex(fullpath,"/i")
	aimspath = string(fullpath)[0:lens]
	aimspath = aimspath+"/cron/"
	//给目录符权限
	err := os.MkdirAll(aimspath, 0777)
	if err != nil {
		seelog.Warn("给缩略图目录设置权限失败")
	}
	//打开图片句柄
	f, err := os.Open(fullpath)
	if err != nil{
		seelog.Warn("图片读取失败")
	}
	defer f.Close()
	//获取图片信息
	img, _, errs := image.Decode(f)
	if errs != nil{
		seelog.Warn("图片读取失败")
	}

	//获取原图宽高
	w,h:=wh(fullpath)

	//按比例截取宽高
	resetW := w / num
	resetH := h / num

	//缩略
	dst :=image.NewNRGBA(image.Rect(0,0,resetW,resetH))
	if graphics.Scale(dst, img) != nil {
		seelog.Warn("图片缩略失败")
	}

    //创建目标保存文件句柄
	imgfile, err := os.Create(aimspath+ "cron_"+ filename)
	defer imgfile.Close()
	//保存图片
	if png.Encode(imgfile,dst) != nil {
		seelog.Warn("缩略图保存失败")
	}
	return aimspath + "cron_"+ filename
}


/**
* 获取图片宽高
 */
func wh (imagePath string) (width,height int) {
	file, err := os.Open(imagePath)
	defer file.Close();
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file) // Image Struct
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	//fmt.Println("Width:", image.Width, "Height:", image.Height)
	return image.Width,image.Height
}