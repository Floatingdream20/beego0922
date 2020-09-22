package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://www.baidu.com"
	c.Data["Email"] = "astaxie@2161979329@qq.com"
	c.TplName = "index.tpl"
}
