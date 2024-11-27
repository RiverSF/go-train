package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	_ "net/http/pprof"
)

//https://gin-gonic.com/zh-cn/docs/

var db = make(map[string]string)

type LoginForm struct {
	User string `json:"user" xml:"user" binding:"required"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// 注册中间件
	//r.Use(MiddleWare())

	//POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	//
	//names[first]=thinkerou&names[second]=tianou
	r.POST("/post", func(c *gin.Context) {
		x := c.Query("x")
		y := c.DefaultQuery("y", "0")
		ids := c.QueryMap("ids")

		z := c.PostForm("names")
		z_def := c.DefaultPostForm("z", "z_def")
		z_arr := c.PostFormArray("names")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{"x": x, "y": y, "z": z, "z_def": z_def, "z_arr": z_arr, "ids": ids, "names": names})
	})

	// 将 request body 绑定到不同的结构体中
	r.POST("/json/map", func(c *gin.Context) {
		objA := LoginForm{}

		// c.ShouldBind 使用了 c.Request.Body，不可重用
		if errA := c.ShouldBind(&objA); errA == nil {
			c.JSON(http.StatusOK, objA)
		}

		// ShouldBindBodyWith 读取 c.Request.Body 并将结果存入上下文。
		// ShouldBindBodyWith 会在绑定之前将 body 存储到上下文中。 这会对性能造成轻微影响，如果调用一次就能完成绑定的话，那就不要用这个方法。
		//只有某些格式需要此功能，如 JSON, XML, MsgPack, ProtoBuf。 对于其他格式, 如 Query, Form, FormPost, FormMultipart 可以多次调用 c.ShouldBind() 而不会造成任任何性能损失
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/v1", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
