package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	UserID int `json:"userID"`
}

func GetUserID(ctx context.Context) (*UserInfo, error) {
	u, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息错误")
	}
	return u, nil
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	v, ok := ctx.Value(userKey).(*UserInfo)
	return v, ok
}
