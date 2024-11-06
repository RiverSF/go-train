package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	//一、基础路由
	beego.Get("/hw", func(ctx *context.Context) {

		//fmt.Println(config.String("httpport"))
		//
		//db, _ := config.GetSection("db")
		//
		//fmt.Println(db["read_host"])

		ctx.Output.Body([]byte("hello world"))
	})

	//二、RESTful Controller 路由
	//beego.Router("/", &controllers.DefController{})
	//beego.Router("/hello", &controllers.DefController{}, "get:Hello")
	//
	////三、自动路由
	//beego.AutoRouter(&controllers.DefController{})

	// /def/login   调用 ObjectController 中的 Login 方法
	// /def/logout  调用 ObjectController 中的 Logout 方法
	// /def/blog/2013/09/12  调用 ObjectController 中的 Blog 方法，参数如下：map[0:2013 1:09 2:12]
}
