package test

import (
	"log"
	"testing"
	"yves-go/config"
	ystruct2 "yves-go/test/entity"
)

func TestSend(t *testing.T) {
	topic := "client_playback_message_publish"
	//msg.Topic = "community_device_report_topic"
	// solution := ystruct2.SolutionMsgParentBody{}
	// externalJSON := solution.GetMessage()
	// msg.Value = sarama.StringEncoder(externalJSON)

	data := ystruct2.GetTimeLineStr()

	config.InitClient("10.202.244.31:9094")
	_ = config.SendMessage(data, topic)

	defer func() {
		if err := config.Close(); err != nil {
			log.Println("close client error", err)
		}
	}()

}
