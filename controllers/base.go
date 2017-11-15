package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

func (this *BaseController) Prepare() {
	//sess := this.StartSession()
	//userLogin := sess.Get("userLogin")
	userLogin := this.GetSession("userLogin")
	fmt.Println("userLogin", userLogin)
	if userLogin == nil {
		this.isLogin = false
	} else {
		this.isLogin = true
	}
	this.Data["isLogin"] = this.isLogin
}

func (this *BaseController) Go404() {
	this.TplName = "404.tpl"
	return
}
