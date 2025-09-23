package interfaces

import (
	"context"
	"yves-go/entity"
)

type UserService interface {
	GetUserById(ctx context.Context, id int64) (*entity.UserBO, error)
	SayHello(ctx context.Context, userName string) (string, error)
}
