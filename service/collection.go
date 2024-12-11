package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/consts"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
)

var collectionServ *CollectionService
var collectionServOnce sync.Once

type CollectionService struct{}

func GetCollectionServ() *CollectionService {
	collectionServOnce.Do(func() {
		collectionServ = &CollectionService{}
	})
	return collectionServ
}

func (c *CollectionService) ShowCollection(ctx *gin.Context, req types.ShowCollectionReq) (resp interface{}, err error) {
	id := ctx.GetInt("id")
	co := dao.NewCollection(ctx)
	collectionList, err := co.FindByID(req, id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	total, err := co.CountAll(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var collection []types.Collection
	for _, v := range collectionList {
		collection = append(collection, types.Collection{
			ID:          v.GoodsID,
			Title:       v.Goods.GoodsName,
			Description: v.Goods.Details,
			Price:       v.Goods.Price,
			ImageUrl:    v.Goods.GoodsImages,
			CreateTime:  v.CreatedTime.Format(consts.TimeFormat),
		})
	}
	var listResp types.CollectionListResp
	listResp.Collections = collection
	listResp.Total = int(total)
	return listResp, nil
}
