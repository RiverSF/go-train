package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type DefController struct {
	beego.Controller
}

func init() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			println("def初始化")
		}
	}()
}

func (c *DefController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Hello struct {
	Data string `json:"data,omitempty"`
}

func (c *DefController) Hello() {
	//c.Ctx.WriteString("hello, world")
	c.Ctx.JSONResp(Hello{"hello, world"})
}

func (c DefController) Hw() {
	//c.Ctx.Output.Body([]byte("Hello, world"))
	c.Ctx.Output.JSON(Hello{"hello, world"}, false, true)
}
