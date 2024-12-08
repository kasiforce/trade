package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
)

var goodsServ *GoodsService
var goodsServOnce sync.Once

type GoodsService struct{}

func GetGoodsService() *GoodsService {
	goodsServOnce.Do(func() {
		goodsServ = &GoodsService{}
	})
	return goodsServ
}

// ShowAllGoods 获取所有商品（管理员端）
func (s *GoodsService) ShowAllGoods(ctx context.Context, req types.ShowAllGoodsReq) (resp interface{}, err error) {
	goods := dao.NewGoods(ctx)
	goodsList, err := goods.AdminFindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	// 创建一个列表来存放最终的返回数据
	var respList []types.GoodsInfo
	for _, goodsInfo := range goodsList {
		respList = append(respList, types.GoodsInfo{
			GoodsID:        goodsInfo.GoodsID,
			GoodsName:      goodsInfo.GoodsName,
			Price:          goodsInfo.Price,
			CategoryName:   goodsInfo.CategoryName,
			Details:        goodsInfo.Details,
			IsSold:         goodsInfo.IsSold,
			GoodsImages:    goodsInfo.GoodsImages,
			CreatedTime:    goodsInfo.CreatedTime,
			UserName:       goodsInfo.UserName,
			Province:       goodsInfo.Province,
			City:           goodsInfo.City,
			District:       goodsInfo.District,
			Address:        goodsInfo.Address,
			Star:           goodsInfo.Star,
			View:           goodsInfo.View,
			DeliveryMethod: goodsInfo.DeliveryMethod,
			ShippingCost:   goodsInfo.ShippingCost,
		})
	}
	// 返回分页后的结果
	var response types.GoodsListResp
	response.ProductList = respList
	response.PageNum = req.PageNum
	response.Total = len(respList)
	return response, nil
}

// 获取已售出商品
func (s *GoodsService) IsSoldGoods(ctx *gin.Context) (resp interface{}, err error) {
	id := ctx.GetInt("id")
	goods := dao.NewGoods(ctx)
	// 直接调用 DAO 层的 IsSoldGoods 方法获取已售出的商品
	filteredGoodsList, err := goods.IsSoldGoods(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 将查询结果转换为接口返回的格式
	var respList []types.GoodsInfo2
	for _, goodsInfo := range filteredGoodsList {
		respList = append(respList, types.GoodsInfo2{
			GoodsID:     goodsInfo.GoodsID,
			GoodsName:   goodsInfo.GoodsName,
			Price:       goodsInfo.Price,
			Details:     goodsInfo.Details,
			GoodsImages: goodsInfo.GoodsImages,
			CreatedTime: goodsInfo.CreatedTime,
		})
	}

	// 返回分页后的结果
	var response types.GoodsListResp2
	response.Data = respList
	return respList, nil
}

// 当前用户获取发布的所有商品
func (s *GoodsService) ShowPublishedGoods(ctx *gin.Context) (resp interface{}, err error) {
	id := ctx.GetInt("id")
	goods := dao.NewGoods(ctx)
	goodsList, err := goods.UserFindAll(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	// 创建一个列表来存放最终的返回数据
	var respList []types.GoodsInfo3
	for _, goodsInfo := range goodsList {
		respList = append(respList, types.GoodsInfo3{
			GoodsID:        goodsInfo.GoodsID,
			GoodsName:      goodsInfo.GoodsName,
			Price:          goodsInfo.Price,
			CategoryName:   goodsInfo.CategoryName,
			Details:        goodsInfo.Details,
			IsSold:         goodsInfo.IsSold,
			GoodsImages:    goodsInfo.GoodsImages,
			CreatedTime:    goodsInfo.CreatedTime,
			UserName:       goodsInfo.UserName,
			Province:       goodsInfo.Province,
			City:           goodsInfo.City,
			District:       goodsInfo.District,
			Star:           goodsInfo.Star,
			View:           goodsInfo.View,
			DeliveryMethod: goodsInfo.DeliveryMethod,
			ShippingCost:   goodsInfo.ShippingCost,
			UserID:         goodsInfo.UserID,
		})
	}
	// 返回分页后的结果
	return respList, nil
}

// DeleteGoods 删除商品
func (s *GoodsService) DeleteGoods(ctx context.Context, id int) (resp interface{}, err error) {
	err = dao.NewGoods(ctx).DeleteGoods(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return resp, nil
	}
	// 创建一个空的返回结构
	resp = map[string]interface{}{}
	return resp, nil
}

// 获取商品列表
func (s *GoodsService) ShowGoodsList(ctx context.Context, req types.ShowGoodsListReq) (resp interface{}, err error) {
	goods := dao.NewGoods(ctx)
	goodsList, err := goods.FindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 创建一个列表来存放最终的返回数据
	var respList []types.GoodsInfo4
	for _, goodsInfo := range goodsList {
		respList = append(respList, types.GoodsInfo4{
			GoodsID:     goodsInfo.GoodsID,
			GoodsName:   goodsInfo.GoodsName,
			Price:       goodsInfo.Price,
			CategoryID:  goodsInfo.CategoryID,
			GoodsImages: goodsInfo.GoodsImages,
		})
	}
	return respList, nil
}

// 更新view
func (s *GoodsService) IncreaseGoodsView(ctx context.Context, goodsID uint) error {
	g := dao.NewGoods(ctx)
	return g.IncreaseView(goodsID)
}

// 当前用户获取发布的所有商品
func (s *GoodsService) FilterGoods(ctx *gin.Context, req types.ShowGoodsReq) (resp interface{}, err error) {
	goods := dao.NewGoods(ctx)
	goodsList, err := goods.FilterGoods(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 创建一个列表来存放最终的返回数据
	var respList []types.GoodsInfo
	for _, goodsInfo := range goodsList {
		respList = append(respList, types.GoodsInfo{
			GoodsID:        goodsInfo.GoodsID,
			GoodsName:      goodsInfo.GoodsName,
			Price:          goodsInfo.Price,
			CategoryName:   goodsInfo.CategoryName,
			Details:        goodsInfo.Details,
			IsSold:         goodsInfo.IsSold,
			GoodsImages:    goodsInfo.GoodsImages,
			CreatedTime:    goodsInfo.CreatedTime,
			UserName:       goodsInfo.UserName,
			Province:       goodsInfo.Province,
			City:           goodsInfo.City,
			District:       goodsInfo.District,
			Address:        goodsInfo.Address,
			Star:           goodsInfo.Star,
			View:           goodsInfo.View,
			DeliveryMethod: goodsInfo.DeliveryMethod,
			ShippingCost:   goodsInfo.ShippingCost,
		})
	}
	// 返回分页后的结果
	return respList, nil
}
