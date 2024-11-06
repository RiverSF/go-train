package main

import (
	_ "api/routers"
	"api/utils/logger"
	"api/utils/signalx"
	"math/rand"
	"runtime"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	if err := logger.Init(); err != nil {
		panic(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())

	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Critical("http server panic err:", err)
			}
		}()

		logger.Debug("beego run...")
		beego.Run()
	}()

	signalx.SignalHandler(serverCloseCallback)
}

func serverCloseCallback() {
	//机器宕机后执行操作
}
