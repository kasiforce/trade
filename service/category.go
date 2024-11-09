package service

import (
	"context"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/types"
	"sync"
)

var CategoryServ *CategoryService
var CategoryServOnce sync.Once

type CategoryService struct {
}

func GetCategoryService() *CategoryService {
	CategoryServOnce.Do(func() {
		CategoryServ = &CategoryService{}
	})
	return CategoryServ
}

func (cs *CategoryService) ShowCategory(ctx context.Context) (resp interface{}, err error) {
	c := dao.NewCategory(ctx)
	categoryList, err := c.FindAll()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.CategoryResp
	for _, ca := range categoryList {
		respList = append(respList, types.CategoryResp{
			CategoryID:   ca.CategoryID,
			CategoryName: ca.CategoryName,
			Descriptions: ca.Description,
		})
	}
	return respList, nil
}
