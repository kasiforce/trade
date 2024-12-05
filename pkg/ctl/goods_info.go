package ctl

import (
	"context"
	"errors"
)

var goodsKey key

type GoodsInfo struct {
	GoodsID int `json:"goodsID"`
}

func GetGoodsID(ctx context.Context) (*GoodsInfo, error) {
	u, ok := FromGoodsContext(ctx)
	if !ok {
		return nil, errors.New("获取商品信息错误")
	}
	return u, nil
}

func NewGoodsContext(ctx context.Context, u *GoodsInfo) context.Context {
	return context.WithValue(ctx, goodsKey, u)
}

func FromGoodsContext(ctx context.Context) (*GoodsInfo, bool) {
	v, ok := ctx.Value(goodsKey).(*GoodsInfo)
	return v, ok
}
