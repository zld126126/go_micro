package service

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"go_micro/pkg/config"
	"go_micro/proto"
)

type QueryService struct {
	Config *config.Config
}

func (q *QueryService) GetUser(ctx context.Context, in *query.UserId, out *query.User) error {
	logrus.Println("get user success")
	out.Name = fmt.Sprint(in.Id) + "哈哈"
	return nil
}

func (q *QueryService) GetActivity(ctx context.Context, in *query.Name, out *query.Activity) error {
	logrus.Println("get act success")
	return nil
}
