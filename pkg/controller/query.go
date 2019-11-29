package controller

import (
	"context"

	"go_micro/pkg/service"
	"go_micro/proto"
)

type Query struct {
	QueryService *service.QueryService
}

func (q *Query) GetUser(ctx context.Context, in *query.UserId, out *query.User) error {
	return q.QueryService.GetUser(ctx, in, out)
}

func (q *Query) GetActivity(ctx context.Context, in *query.Name, out *query.Activity) error {
	return q.QueryService.GetActivity(ctx, in, out)
}
