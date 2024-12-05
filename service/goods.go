package service

import (
	"context"
	"github.com/cghdjvjg/trade/pkg/ctl"
	"github.com/cghdjvjg/trade/pkg/util"
	"github.com/cghdjvjg/trade/repository/db/dao"
	"github.com/cghdjvjg/trade/types"
	"strconv"
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
			GoodsID:      goodsInfo.GoodsID,
			GoodsName:    goodsInfo.GoodsName,
			Price:        goodsInfo.Price,
			CategoryName: goodsInfo.CategoryName,
			Details:      goodsInfo.Details,
			IsSold:       goodsInfo.IsSold,
			GoodsImages:  goodsInfo.GoodsImages,
			CreatedTime:  goodsInfo.CreatedTime,
			UserName:     goodsInfo.UserName,
			Province:     goodsInfo.Province,
			City:         goodsInfo.City,
			District:     goodsInfo.District,
			Star:         goodsInfo.Star,
			View:         goodsInfo.View,
			PayMethod:    strconv.Itoa(goodsInfo.PayMethod),
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
func (s *GoodsService) IsSoldGoods(ctx context.Context, req types.IsSoldGoodsResp) (resp interface{}, err error) {
	goods := dao.NewGoods(ctx)
	// 直接调用 DAO 层的 IsSoldGoods 方法获取已售出的商品
	filteredGoodsList, err := goods.IsSoldGoods(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	// 将查询结果转换为接口返回的格式
	var respList []types.GoodsInfo2
	for _, goodsInfo := range filteredGoodsList {
		respList = append(respList, types.GoodsInfo2{
			GoodsID:   goodsInfo.GoodsID,
			GoodsName: goodsInfo.GoodsName,
			UserID:    goodsInfo.UserID,
			Price:     goodsInfo.Price,
			//CategoryID:  goodsInfo.CategoryID,
			Details: goodsInfo.Details,
			//IsSold:      goodsInfo.IsSold,
			GoodsImages: goodsInfo.GoodsImages,
			CreatedTime: goodsInfo.CreatedTime,
		})
	}

	// 返回分页后的结果
	var response types.GoodsListResp2
	response.Data = respList
	return respList, nil
}

/*
// ShowGoodsDetail 获取商品详情
func (s *GoodsService) ShowGoodsDetail(ctx context.Context, id int) (resp interface{}, err error) {
	goods := dao.NewGoods(ctx)
	goodsInfo, err := goods.FindByID(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = types.GoodsDetailResp{
		GoodsID:     goodsInfo.GoodsID,
		Name:        goodsInfo.Name,
		Category:    goodsInfo.Category,
		Price:       goodsInfo.Price,
		Stock:       goodsInfo.Stock,
		IsSold:      goodsInfo.IsSold,
		Picture:     goodsInfo.Picture,
		CreatedAt:   goodsInfo.CreatedAt,
		Description: goodsInfo.Description,
	}
	return
}

/*
// FilterGoods 按条件筛选商品
func (s *GoodsService) FilterGoods(ctx context.Context, filter map[string]interface{}, req types.ShowGoodsReq) (resp interface{}, err error) {
	goods := dao.NewGoods(ctx)
	filteredGoodsList, err := goods.Filter(filter, req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.GoodsInfo
	for _, goodsInfo := range filteredGoodsList {
		respList = append(respList, types.GoodsInfo{
			GoodsID:   goodsInfo.GoodsID,
			Name:      goodsInfo.Name,
			Category:  goodsInfo.Category,
			Price:     goodsInfo.Price,
			Stock:     goodsInfo.Stock,
			IsSold:    goodsInfo.IsSold,
			Picture:   goodsInfo.Picture,
			CreatedAt: goodsInfo.CreatedAt,
		})
	}
	var response types.GoodsListResp
	response.GoodsList = respList
	response.PageNum = req.PageNum
	response.Total = len(respList)
	return response, nil
}

// CreateGoods 创建商品
func (s *GoodsService) CreateGoods(ctx context.Context, req types.GoodsInfo) (resp interface{}, err error) {
	if req.Name == "" || req.Price <= 0 || req.Stock < 0 || req.Category == "" {
		err = errors.New("参数不能为空")
		return
	}
	goods := dao.NewGoods(ctx)
	modelGoods := map[string]interface{}{
		"name":        req.Name,
		"category":    req.Category,
		"price":       req.Price,
		"stock":       req.Stock,
		"picture":     req.Picture,
		"description": req.Description,
		"isSold":      req.IsSold,
	}
	err = goods.CreateGoods(modelGoods)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}
*/
// DeleteGoods 删除商品
func (s *GoodsService) DeleteGoods(ctx context.Context) (resp interface{}, err error) {
	goods, err := ctl.GetGoodsID(ctx)
	if err != nil {
		util.LogrusObj.Error(err)
		return resp, nil
	}
	a := dao.NewGoods(ctx)
	err = a.DeleteGoods(goods.GoodsID)
	if err != nil {
		util.LogrusObj.Error(err)
		return resp, nil
	}
	// 创建一个空的返回结构
	resp = map[string]interface{}{}
	return resp, nil
}
