package task

import (
	"time"
	"yves-go/util"

	"github.com/apolloconfig/agollo/v4/component/log"
)

func AddWeatherTask() {
	ticket := time.NewTicker(time.Second * 30)
	defer ticket.Stop()

	go func() {
		for {
			select {
			case <-ticket.C:
				log.Info("AddWeatherTask start")
				AddWeather2Cache()
			}
		}
	}()

}

func AddWeather2Cache() {
	log.Info("AddWeather2Cache start id:", util.NewSnowflake())
}
