package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type Project7Controller struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/
	c.Data["content"] = "this is jade template"
	c.TplName = "home.jade"
}

func (c *Project7Controller) Get() {

	c.TplName = "Project7-Yellow-Studio.jade"
}
