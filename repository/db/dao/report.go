package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type Report struct {
	*gorm.DB
}

func NewReportByDB(db *gorm.DB) *Report {
	return &Report{db}
}

func NewReport(ctx context.Context) *Report {
	return &Report{NewDBClient(ctx)}
}

func (r *Report) FindAll() (reports []*model.Report, err error) {
	err = r.DB.Model(&model.Report{}).Find(&reports).Error
	return
}

func (r *Report) FindByID(id int) (report *model.Report, err error) {
	err = r.DB.Model(&model.Report{}).Where("reportID = ?", id).First(&report).Error
	return
}

func (r *Report) CreateReport(report *model.Report) (err error) {
	err = r.DB.Model(&model.Report{}).Create(&report).Error
	return
}

func (r *Report) UpdateReport(id int, report *model.Report) (err error) {
	err = r.DB.Model(&model.Report{}).Where("reportID = ?", id).Updates(report).Error
	return
}

func (r *Report) DeleteReport(id int) (err error) {
	err = r.DB.Model(&model.Report{}).Where("reportID = ?", id).Delete(&model.Report{}).Error
	return
}
