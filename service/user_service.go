package service

import (
	"context"
	"log"
	"strconv"
	"yves-go/entity"
	"yves-go/interfaces"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func DubboInit() {
	log.Println("user_service init")

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Dubbo初始化遇到问题，但不影响项目启动: %v", r)
		}
	}()

	// 先注册服务提供者
	config.SetProviderService(&UserServiceImpl{})
	// 然后加载配置
	err := config.Load(config.WithPath("./conf/dubbogo.yaml"))
	if err != nil {
		log.Printf("Dubbo config load error: %v", err)
		return
	}

	log.Println("Dubbo config loaded successfully")
	log.Printf("Config: %+v\n", config.GetProviderConfig())

	log.Println("Dubbo service started and registered to ZooKeeper")
}

type UserServiceImpl struct {
}

// 确保实现了interfaces.UserService接口
var _ interfaces.UserService = (*UserServiceImpl)(nil)

func (u *UserServiceImpl) GetUserById(ctx context.Context, id int64) (*entity.UserBO, error) {
	user := &entity.UserBO{
		Id:       id,
		UserName: "User:" + strconv.FormatInt(id, 10),
	}
	return user, nil
}

func (u *UserServiceImpl) SayHello(ctx context.Context, userName string) (string, error) {
	return "Hello " + userName, nil
}
