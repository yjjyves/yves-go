package config

import (
	"fmt"

	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer

func InitClient(address string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true

	var err error
	client, err = sarama.NewSyncProducer([]string{address}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
}

func SendMessage(data string, topic string) error {
	msg := &sarama.ProducerMessage{}
	msg.Value = sarama.StringEncoder(data)
	msg.Topic = topic
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("pid:%v offset:%v\n", partition, offset)
	return nil
}

func Close() error {
	if client == nil {
		return nil // 如果client为nil，不需要关闭
	}
	return client.Close()
}
