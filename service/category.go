package service

import (
	"context"
	"errors"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/repository/db/model"
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

func (cs *CategoryService) ShowCategory(ctx context.Context, req types.ShowCategoryReq) (resp interface{}, err error) {
	c := dao.NewCategory(ctx)
	categoryList, err := c.FindAll(req)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	total, err := c.CountAll()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.CategoryInfo
	for _, ca := range categoryList {
		respList = append(respList, types.CategoryInfo{
			CategoryID:   ca.CategoryID,
			CategoryName: ca.CategoryName,
			Descriptions: ca.Description,
		})
	}
	var response types.CategoryList
	response.Categories = respList
	response.PageNum = req.PageNum
	response.Total = int(total)
	return response, nil
}
func (cs *CategoryService) ShowUserCategory(ctx context.Context) (resp interface{}, err error) {
	c := dao.NewCategory(ctx)
	categoryList, err := c.ShowCategory()
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.CategoryInfo
	for _, ca := range categoryList {
		respList = append(respList, types.CategoryInfo{
			CategoryID:   ca.CategoryID,
			CategoryName: ca.CategoryName,
			Descriptions: ca.Description,
		})
	}
	return respList, nil
}
func (cs *CategoryService) AddCategory(ctx context.Context, req types.CategoryInfo) (resp interface{}, err error) {
	if req.CategoryName == "" {
		err = errors.New("类别名称不能为空")
		return
	}
	c := dao.NewCategory(ctx)
	modelCategory := &model.Category{
		CategoryName: req.CategoryName,
		Description:  req.Descriptions,
	}
	uid, err := c.AddCategory(modelCategory)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = types.CategoryID{
		CategoryID: uid,
	}
	return
}
func (cs *CategoryService) UpdateCategory(ctx context.Context, req types.CategoryInfo) (resp interface{}, err error) {
	c := dao.NewCategory(ctx)
	modelCategory := &model.Category{
		CategoryID:   req.CategoryID,
		CategoryName: req.CategoryName,
		Description:  req.Descriptions,
	}
	err = c.UpdateCategory(req.CategoryID, modelCategory)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, id int) (resp interface{}, err error) {
	c := dao.NewCategory(ctx)
	err = c.DeleteCategory(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}
