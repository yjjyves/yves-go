package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"yves-go/req"
	"yves-go/service"

	"github.com/gin-gonic/gin"
)

func LogTime1(c *gin.Context) {
	t := time.Now()
	fmt.Println("中间件开始执行了")
	// 设置变量到Context的key中，可以通过Get()取
	c.Set("request", "中间件")
	status := c.Writer.Status()
	fmt.Println("中间件执行完毕", status)
	t2 := time.Since(t)
	fmt.Println("time:", t2)
}

func LogTime2(c *gin.Context) {
	t := time.Now()
	fmt.Println("中间件开始执行了")
	// 设置变量到Context的key中，可以通过Get()取
	c.Set("request", "中间件")
	// 执行函数
	c.Next()
	// 中间件执行完后续的一些事情
	status := c.Writer.Status()
	fmt.Println("中间件执行完毕", status)
	t2 := time.Since(t)
	fmt.Println("time:", t2)
}

func RegisterNewsQueryApi(r *gin.Engine) {
	log.Println("RegisterNewsQueryApi /news/query")
	r.POST("/news/query", LogTime2, func(c *gin.Context) {
		// 声明接收的变量
		var queryReq req.NewsQueryReqVO

		// 解析 JSON
		if err := c.ShouldBindJSON(&queryReq); err != nil {
			// 返回错误信息
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "400",
				"error":  err.Error(),
			})
			return
		}

		// 解析成功，调用服务
		result := service.EsNewsSearchServiceInstance.EsNewsSearch(&queryReq)
		// 返回结果
		c.JSON(http.StatusOK, gin.H{
			"status": "200",
			"data":   result,
		})
	})
}

func RegisterUserQueryApi(r *gin.Engine) {
	log.Println("RegisterUserQueryApi /news/query")
	r.POST("/user/query/:id", LogTime2, func(c *gin.Context) {
		// 声明接收的变量
		var userId = c.Param("id")
		log.Println("userId:", userId)
		// 解析成功，调用服务
		// 返回结果
		c.JSON(http.StatusOK, gin.H{
			"status": "200",
			"data":   "result",
		})
	})
}
