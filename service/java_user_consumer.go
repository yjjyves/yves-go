package service

import (
	"context"
	"log"
	"yves-go/interfaces"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var JavaUserServiceClient *JavaUserServiceImpl

func InitJavaUserServiceConsumer() {
	log.Println("初始化Java UserService消费者...")

	// 设置消费者服务引用 - 这是关键步骤
	config.SetConsumerService(&JavaUserServiceImpl{})

	log.Println("Java UserService消费者初始化完成")
}

// JavaUserServiceImpl 消费者存根类 - 按照官方文档的方式实现
type JavaUserServiceImpl struct {
	// 使用函数字段，Dubbo会自动代理到远程Java服务
	GetUserById func(ctx context.Context, id int64) (*interfaces.User, error) `dubbo:"getUserById"`
	SayHello    func(ctx context.Context, name string) (string, error)        `dubbo:"sayHello"`
}

func GetJavaUserServiceClient() *JavaUserServiceImpl {
	if JavaUserServiceClient == nil {
		log.Println("=== 开始获取Java服务客户端 ===")

		// 获取消费者引用 - 使用正确的key "JavaUserServiceImpl"
		client := config.GetConsumerService("JavaUserServiceImpl")
		log.Printf("config.GetConsumerService('JavaUserServiceImpl') 返回: %v", client)

		if client != nil {
			log.Printf("客户端类型: %T", client)

			// 尝试类型转换
			if javaClient, ok := client.(*JavaUserServiceImpl); ok {
				JavaUserServiceClient = javaClient
				log.Println("✓ 成功获取远程Java服务客户端")
				log.Printf("✓ 客户端类型转换成功: %T", JavaUserServiceClient)
			} else {
				log.Printf("✗ 类型转换失败，期望类型: *JavaUserServiceImpl，实际类型: %T", client)
				return nil
			}
		} else {
			log.Printf("✗ 无法获取远程Java服务，请检查配置和服务注册")
			log.Println("=== 检查Dubbo配置 ===")

			// 检查Dubbo配置
			providerConfig := config.GetProviderConfig()
			log.Printf("Provider配置: %+v", providerConfig)

			consumerConfig := config.GetConsumerConfig()
			log.Printf("Consumer配置: %+v", consumerConfig)

			return nil
		}
	} else {
		log.Println("✓ 使用已缓存的Java服务客户端")
	}

	log.Printf("=== 返回Java服务客户端: %v ===", JavaUserServiceClient)
	return JavaUserServiceClient
}

// 调用Java服务的示例方法
func CallJavaUserService() {
	log.Println("=== 开始调用Java服务 ===")

	client := GetJavaUserServiceClient()
	if client == nil {
		log.Println("✗ Java UserService客户端未初始化")
		return
	}

	log.Println("✓ Java UserService客户端已获取")
	ctx := context.Background()

	// 调用getUserById方法
	log.Println("=== 调用getUserById方法 ===")
	log.Printf("调用参数: id=123")

	// 直接调用函数字段
	user, err := client.GetUserById(ctx, 123)
	if err != nil {
		log.Printf("✗ 调用getUserById失败: %v", err)
	} else {
		log.Printf("✓ 获取到用户: %+v", user)
		log.Printf("✓ 用户类型: %T", user)
	}

	// 调用sayHello方法
	log.Println("=== 调用sayHello方法 ===")
	log.Printf("调用参数: name='Go客户端'")

	hello, err := client.SayHello(ctx, "Go客户端")
	if err != nil {
		log.Printf("✗ 调用sayHello失败: %v", err)
	} else {
		log.Printf("✓ Java服务返回: %s", hello)
		log.Printf("✓ 返回类型: %T", hello)
	}

	log.Println("=== Java服务调用完成 ===")
}
