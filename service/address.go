package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/repository/db/model"
	"github.com/kasiforce/trade/types"
	"sync"
)

var AddrService *AddressService
var AddrServOnce sync.Once

type AddressService struct{}

func GetAddrService() *AddressService {
	AddrServOnce.Do(func() {
		AddrService = &AddressService{}
	})
	return AddrService
}

func (a *AddressService) AddAddr(c *gin.Context, req types.AddrAddReq) (resp interface{}, err error) {
	modelAddr := &model.Address{
		UserID:    c.GetInt("id"),
		Name:      req.Name,
		Province:  req.Province,
		City:      req.City,
		District:  req.District,
		Address:   req.Address,
		Tel:       req.Tel,
		IsDefault: req.IsDefault,
	}
	ad := dao.NewAddress(c)
	id, err := ad.AddAddr(modelAddr)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	resp = types.AddrAddResp{
		ID: id,
	}
	return
}

func (a *AddressService) ShowAddr(c *gin.Context) (resp interface{}, err error) {
	ad := dao.NewAddress(c)
	addrList, err := ad.GetAddrByID(c.GetInt("id"))
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	var respList []types.AddrInfoResp
	for _, addrInfo := range addrList {
		respList = append(respList, types.AddrInfoResp{
			ID:        addrInfo.ID,
			UserID:    addrInfo.UserID,
			Name:      addrInfo.Name,
			Province:  addrInfo.Province,
			City:      addrInfo.City,
			District:  addrInfo.District,
			Address:   addrInfo.Address,
			Tel:       addrInfo.Tel,
			IsDefault: addrInfo.IsDefault,
		})
	}
	return respList, nil
}

func (a *AddressService) UpdateAddr(c *gin.Context, req types.AddrUpdateReq) (resp interface{}, err error) {
	ad := dao.NewAddress(c)
	modelAddr := &model.Address{
		Name:      req.Name,
		Province:  req.Province,
		City:      req.City,
		District:  req.District,
		Address:   req.Address,
		Tel:       req.Tel,
		IsDefault: req.IsDefault,
	}
	err = ad.UpdateAddr(req.ID, modelAddr)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (a *AddressService) UpdateDefault(c *gin.Context, req types.UpdateDefaultReq) (resp interface{}, err error) {
	ad := dao.NewAddress(c)
	err = ad.UpdateDefaultAddr(req.OldDefault, req.NewDefault)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}

func (a *AddressService) DeleteAddr(c *gin.Context, id int) (resp interface{}, err error) {
	ad := dao.NewAddress(c)
	err = ad.DeleteAddr(id)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	return
}
