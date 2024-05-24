package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
	_ "web/routers"
)

func init() {
	go func() {
		time.Sleep(1)
		println("启动初始化")
	}()
}

func main() {
	beego.Run()
}

