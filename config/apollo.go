package config

import (
	"fmt"
	"log"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

var ApolloClient agollo.Client

type ChangeListener struct{}

func (c *ChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	fmt.Printf("Apollo 配置更新: %+v\n", changeEvent.Changes)
}

func (c *ChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	// 最新配置事件
	fmt.Println("Apollo 最新配置")
}

func init() {
	log.Println("init apollo.....")
	loadAppConfig := &config.AppConfig{
		AppID:          "solution-smart-monitor",
		Cluster:        "default",
		NamespaceName:  "application",
		IP:             "http://10.202.244.23:32002/config",
		IsBackupConfig: true,
	}

	ApolloClient, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return loadAppConfig, nil
	})
	if err != nil {
		log.Fatalf("start agollo error: %s", err)
	}
	ApolloClient.AddChangeListener(&ChangeListener{})

	// 读取配置示例
	go func() {
		for {
			value := ApolloClient.GetConfig(loadAppConfig.NamespaceName).GetValue("zookeeper.address")
			fmt.Println("Apollo 配置 zookeeper.address =", value)
			time.Sleep(5 * time.Second)
		}
	}()
}
