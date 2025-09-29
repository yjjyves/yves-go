package task

import (
	"time"
	"yves-go/util"

	"github.com/sirupsen/logrus"
)

func AddWeatherTask() {
	logrus.Info("AddWeatherTask init")
	ticket := time.NewTicker(time.Second * 10)

	go func() {
		for {
			select {
			case <-ticket.C:
				logrus.Info("AddWeatherTask start")
				AddWeather2Cache()
			}
		}
	}()
}

func AddWeather2Cache() {
	pool := NewWorkerPool(2, 10)
	pool.Submit(func() {
		logrus.Info("AddWeather2Cache start1 id:", util.NewSnowflake())
		time.Sleep(time.Second * 2)
	})

	pool.Submit(func() {
		logrus.Info("AddWeather2Cache start2 id:", util.NewSnowflake())
		time.Sleep(time.Second * 2)
	})
}
