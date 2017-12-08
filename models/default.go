package models

import (
	//"github.com/astaxie/beego/orm"
	//"time"
	"time"
)

/**
 * 文件描述
 * Create on 11/29/17 10:57 AM
 * Create by hu
 */
type Article struct{
	Id 		int  `orm:"auto";form:"-"`
	Title 	string	`form:"title"`
	Img 	string	`form:"img"`
	Content string	 `form:"content"`
	//CreateTime string	 `orm:"type(datetime)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);null;auto_now_add"`
	//Created time.Time `orm:"auto_now_add;type(datetime)"`
}