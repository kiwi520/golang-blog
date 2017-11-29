package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
/**
 * 文件描述
 * Create on 11/29/17 10:57 AM
 * Create by hu
 */
type Article struct{
	Id 		int
	Title 	string
	content string
}

func init()  {
	// register model
	orm.RegisterModel(new(Article))

	// create table
	orm.RunSyncdb("default", false, true)
}