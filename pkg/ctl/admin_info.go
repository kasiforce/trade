package ctl

import (
	"context"
	"errors"
)

var adminKey key

type AdminInfo struct {
	AdminID int `json:"adminID"`
}

func GetAdminID(ctx context.Context) (*AdminInfo, error) {
	u, ok := FromAdminContext(ctx)
	if !ok {
		return nil, errors.New("获取管理员信息错误")
	}
	return u, nil
}

func NewAdminContext(ctx context.Context, u *AdminInfo) context.Context {
	return context.WithValue(ctx, adminKey, u)
}

func FromAdminContext(ctx context.Context) (*AdminInfo, bool) {
	v, ok := ctx.Value(adminKey).(*AdminInfo)
	return v, ok
}
