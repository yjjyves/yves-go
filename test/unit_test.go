package test

import (
	"fmt"
	"testing"
	"time"
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
