package controllers

import (
	"github.com/astaxie/beego"
)

type FinanceController struct {
	beego.Controller
}

func (this *FinanceController) Get() {
	m := make(map[string]interface{})
	m["id"] = 12
	m["name"] = "pin"
	this.Data["json"] = m
	this.ServeJSON()
}
