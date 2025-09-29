package main

import (
	"fmt"
	"net/http"
	"os"
	"yves-go/api"
	"yves-go/service"
	"yves-go/task"
	"yves-go/video"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	api.RegisterNewsQueryApi(r)
	api.RegisterUserQueryApi(r)
	video.RegisterVideoOffer(r)
	task.AddWeatherTask()

	// 先初始化Java服务消费者（必须在DubboInit之前）
	//service.InitJavaUserServiceConsumer()
	//
	//// 然后初始化Dubbo服务
	//service.DubboInit()

	// 添加测试Java服务的API
	r.GET("/test-java-service", func(c *gin.Context) {
		service.CallJavaUserService()
		c.JSON(http.StatusOK, gin.H{
			"message": "Java服务调用完成，请查看日志",
		})
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err := r.Run(":8001")
	if err != nil {
		fmt.Println("启动失败!")
	}

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}
