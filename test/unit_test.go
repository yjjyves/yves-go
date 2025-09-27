package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func waitReceive(chan1 chan int) {
	chan2 := make(chan int)
	for {
		select {
		case r := <-chan1:
			// 如果chan1成功读到数据，则进行该case处理语句
			fmt.Println(r)
		case chan2 <- 1:
			// 如果成功向chan2写入数据，则进行该case处理语句
			fmt.Println("write")
		default:
			// 如果上面都没有成功，则进入default处理流程
		}
	}
}

func TestT1(t *testing.T) {
	chan1 := make(chan int, 10)
	go waitReceive(chan1)
	for i := 0; i < 10; i++ {
		chan1 <- i
	}

	time.Sleep(10 * time.Second)
}

func write(chan1 chan int) {
	for i := 0; i < 10; i++ {
		chan1 <- i
	}
	close(chan1)
}
func TestT2(t *testing.T) {
	chan1 := make(chan int, 10)
	go write(chan1)
	for i := range chan1 {
		fmt.Println(i)
	}

}

func Test3(t *testing.T) {
	// 不同级别日志
	logrus.Debug("调试信息") // 不会输出，因为级别是 Info
	logrus.Info("普通信息")
	logrus.Warn("警告信息")
	logrus.Error("错误信息")
	// logrus.Fatal("致命错误")  // 会退出程序
	// logrus.Panic("Panic日志") // 会 panic

	// 带字段的结构化日志
	logrus.WithFields(logrus.Fields{
		"user_id": 12345,
		"ip":      "192.168.1.1",
		"event":   "login",
	}).Info("用户登录")

	// 更简洁的字段写法
	logrus.
		WithField("user_id", 12345).
		WithField("action", "purchase").
		Info("用户购买商品")

}
