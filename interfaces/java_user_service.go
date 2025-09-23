package interfaces

import (
	"context"
	"log"

	hessian "github.com/apache/dubbo-go-hessian2"
)

// User 对应Java的User类
type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// JavaClassName 返回Java类名，用于序列化
func (u *User) JavaClassName() string {
	return "com.yves.dubbo3.User"
}

// JavaUserService 对应Java的UserService接口
// JavaUserService 接口 - 用于类型断言
type JavaUserService interface {
	GetUserById(ctx context.Context, id int64) (*User, error)
	SayHello(ctx context.Context, name string) (string, error)
}

func init() {
	// 注册POJO到hessian库
	hessian.RegisterPOJO(&User{})
	log.Println("注册User POJO到hessian库")
}
