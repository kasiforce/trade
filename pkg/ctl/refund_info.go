package ctl

import (
	"context"
	"errors"
)

var refundKey key

type RefundInfo struct {
	RefundID int `json:"refundID"`
}

func GetRefundID(ctx context.Context) (*RefundInfo, error) {
	u, ok := FromRefundContext(ctx)
	if !ok {
		return nil, errors.New("获取信息错误")
	}
	return u, nil
}

func NewRefundContext(ctx context.Context, u *RefundInfo) context.Context {
	return context.WithValue(ctx, refundKey, u)
}

func FromRefundContext(ctx context.Context) (*RefundInfo, bool) {
	v, ok := ctx.Value(refundKey).(*RefundInfo)
	return v, ok
}
