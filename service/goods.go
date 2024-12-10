package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
	"time"
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
	// 获取数据库中所有商品的总数
	total, err := goods.CountAll()
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	goodsList, err := goods.AdminFindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	// 创建一个列表来存放最终的返回数据
	var respList []types.GoodsInfo
	for _, goodsInfo := range goodsList {
		// 将 DeliveryMethod 从 int 转换为 string
		var deliveryMethod string
		switch goodsInfo.DeliveryMethod {
		case 0:
			deliveryMethod = "无需快递"
		case 1:
			deliveryMethod = "自提"
		case 2:
			deliveryMethod = "邮寄"
		default:
			deliveryMethod = "未知"
		}
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
			DeliveryMethod: deliveryMethod,
			ShippingCost:   goodsInfo.ShippingCost,
		})
	}
	// 返回分页后的结果
	var response types.GoodsListResp
	response.ProductList = respList
	response.PageNum = req.PageNum
	response.Total = int(total)
	return response, nil
}

// 获取已售出商品
func (s *GoodsService) IsSoldGoods(ctx *gin.Context, id int) (resp interface{}, err error) {
	//id := ctx.GetInt("id")
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
func (s *GoodsService) ShowPublishedGoods(ctx *gin.Context, id int) (resp interface{}, err error) {
	//id := ctx.GetInt("id")
	goods := dao.NewGoods(ctx)
	goodsList, err := goods.UserFindAll(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}

	// 创建一个列表来存放最终的返回数据
	var respList []types.GoodsInfo3
	for _, goodsInfo := range goodsList {
		// 将 DeliveryMethod 从 int 转换为 string
		var deliveryMethod string
		switch goodsInfo.DeliveryMethod {
		case 0:
			deliveryMethod = "无需快递"
		case 1:
			deliveryMethod = "自提"
		case 2:
			deliveryMethod = "邮寄"
		default:
			deliveryMethod = "未知"
		}
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
			Address:        goodsInfo.Address,
			Star:           goodsInfo.Star,
			View:           goodsInfo.View,
			DeliveryMethod: deliveryMethod,
			ShippingCost:   goodsInfo.ShippingCost,
			UserID:         goodsInfo.UserID,
			AddrID:         goodsInfo.AddrID,
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
func (s *GoodsService) IncreaseGoodsView(ctx context.Context, req types.ShowDetailReq) error {
	g := dao.NewGoods(ctx)
	return g.IncreaseView(req)
}

// 筛选商品
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
		// 将 DeliveryMethod 从 int 转换为 string
		var deliveryMethod string
		switch goodsInfo.DeliveryMethod {
		case 0:
			deliveryMethod = "无需快递"
		case 1:
			deliveryMethod = "自提"
		case 2:
			deliveryMethod = "邮寄"
		default:
			deliveryMethod = "未知"
		}
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
			DeliveryMethod: deliveryMethod,
			ShippingCost:   goodsInfo.ShippingCost,
		})
	}
	// 返回分页后的结果
	return respList, nil
}

// 获取商品详情
func (s *GoodsService) ShowGoodsDetail(ctx *gin.Context, req types.ShowDetailReq) (resp interface{}, err error) {
	userid := ctx.GetInt("id")
	goods := dao.NewGoods(ctx)
	goodsInfo, err := goods.ShowGoodsDetail(req, userid)
	if err != nil {
		util.LogrusObj.Error(err)
		return nil, err
	}
	// 将 DeliveryMethod 从 int 转换为 string
	var deliveryMethod string
	switch goodsInfo.DeliveryMethod {
	case 0:
		deliveryMethod = "无需快递"
	case 1:
		deliveryMethod = "自提"
	case 2:
		deliveryMethod = "邮寄"
	default:
		deliveryMethod = "未知"
	}
	// 创建最终的返回数据
	respData := types.ShowGoodsDetail{
		GoodsID:        goodsInfo.GoodsID,
		GoodsName:      goodsInfo.GoodsName,
		Price:          goodsInfo.Price,
		CategoryName:   goodsInfo.CategoryName,
		Details:        goodsInfo.Details,
		IsSold:         goodsInfo.IsSold,
		GoodsImages:    goodsInfo.GoodsImages,
		CreatedTime:    goodsInfo.CreatedTime,
		UserName:       goodsInfo.UserName,
		UserID:         goodsInfo.UserID,
		Province:       goodsInfo.Province,
		City:           goodsInfo.City,
		District:       goodsInfo.District,
		Address:        goodsInfo.Address,
		Star:           goodsInfo.Star,
		View:           goodsInfo.View,
		DeliveryMethod: deliveryMethod,
		ShippingCost:   goodsInfo.ShippingCost,
		AddrID:         goodsInfo.AddrID,
		Tel:            goodsInfo.Tel,
		IsStarred:      goodsInfo.IsStarred,
	}
	// 返回结果
	return respData, nil
}

// 发布闲置
func (s *GoodsService) AddGoods(ctx *gin.Context, req types.CreateGoodsReq) (resp interface{}, err error) {
	userid := ctx.GetInt("id")
	goods := dao.NewGoods(ctx)
	goodsID, err := goods.CreateGoods(req, userid)
	if req.GoodsName == "" || req.Details == "" || req.DeliveryMethod == "" || req.CategoryName == "" || req.Price == 0 {
		err = errors.New("参数不能为空")
		return nil, err
	}
	resp = map[string]interface{}{
		"id": goodsID,
	}
	return resp, nil
}

// 更新商品收藏情况
func (s *GoodsService) UpdateGoodsIsStarred(ctx *gin.Context, goodsID int, r types.IsStarred) (resp interface{}, err error) {
	userid := ctx.GetInt("id") // 获取当前用户的 ID
	isStarred := r.IsStarred   // 获取传递过来的收藏状态
	// 打印 goodsID 和 userid
	// 使用 dao.NewGoods(ctx) 初始化 goods 实例
	goods := dao.NewGoods(ctx)
	// 查询当前商品是否已被用户收藏
	isStarredInDB, err := goods.CheckGoodsIsStarred(goodsID, userid)
	if err != nil {
		return nil, err
	}
	// 如果收藏状态没有变化，则返回“未修改商品收藏详情”
	if isStarred == isStarredInDB {
		resp = map[string]interface{}{
			"data": "未修改商品收藏详情",
		}
		return resp, nil
	}
	// 如果 isStarred 是 true，表示用户希望收藏该商品
	if isStarred {
		err = goods.AddToCollection(goodsID, userid, time.Now()) // 添加到收藏
		if err != nil {
			return nil, err
		}
	} else {
		// 如果 isStarred 是 false，表示用户希望取消收藏
		err = goods.RemoveFromCollection(goodsID, userid) // 从收藏中移除
		if err != nil {
			return nil, err
		}
	}
	// 返回成功响应
	resp = map[string]interface{}{
		"userid": userid,
		"data":   "商品收藏状态已更新",
	}
	return resp, nil
}
