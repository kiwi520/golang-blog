package models

import (
	"time"
	"github.com/astaxie/beego/validation"
	"strings"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/scrypt"
	"encoding/hex"
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
	Content string	 `orm:"column(content);type(text);"form:"content"`
	//CreateTime string	 `orm:"type(datetime)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);null;auto_now_add"`
	//Created time.Time `orm:"auto_now_add;type(datetime)"`
}

type User struct{
	Id 		int  `orm:"auto";form:"-"`
	Email 	string	`form:"email";valid:"Required;Match(/^Bee.*/);Range(1, 40)"`
	Password 	string	`form:"password";valid:"Email; MaxSize(100)"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);null;auto_now_add";form:"-"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *User) Valid(v *validation.Validation) {
	o := orm.NewOrm()
	var user User
	// 获取 QuerySeter 对象，user 为表名
	ps := o.QueryTable("user")
	if len(u.Email) <=0 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("email", "邮箱不能为空")
	}
	if len(u.Email) >0 {

		emailerr :=ps.Filter("email",u.Email).One(&user)

		if emailerr == orm.ErrMultiRows {
			// 多条的时候报错
			v.SetError("email", "该用户数据有重复,请注意")
		}else if emailerr == orm.ErrNoRows {
			// 没有找到记录
			v.SetError("email", "没有该用户,请联系管理员")
		}
	}
	if len(u.Password) ==0 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Password", "密码不能为空")

	}

	if len(u.Password) <8 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Password", "密码至少需要8位")
	}

	if len(u.Password) >=8 {
		dbpasswd := user.Password
		crupasswd := u.Password
		strpasswd,_:= scrypt.Key([]byte(crupasswd), []byte("kiwi"), 16384, 8, 1, 32)
		if hex.EncodeToString(strpasswd) !=dbpasswd {
			v.SetError("Password", "密码不正确,请重试")
		}
	}
	if strings.Index(u.Email, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Email", "邮箱名称里不能含有 admin")
	}
}