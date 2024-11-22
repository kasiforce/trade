package dao

import (
	"context"
	"github.com/kasiforce/trade/repository/db/model"
	"gorm.io/gorm"
)

type School struct {
	*gorm.DB
}

func NewSchoolByDB(db *gorm.DB) *School {
	return &School{db}
}

func NewSchool(ctx context.Context) *School {
	return &School{NewDBClient(ctx)}
}

func (s *School) FindSchoolID(name string) (id int, err error) {
	school := &model.School{}
	err = s.DB.Model(&model.School{}).Where("schoolName = ?", name).First(&school).Error
	id = school.SchoolID
	return
}
func (s *School) FindAll() (schools []*model.School, err error) {
	err = s.DB.Model(&model.School{}).Find(&schools).Error
	return
}

func (s *School) FindByID(id int) (sch *model.School, err error) {
	err = s.DB.Model(&model.School{}).Where("schoolID = ?", id).First(&sch).Error
	return
}

func (s *School) FindByName(name string) (exist bool, err error) {
	var cnt int64
	err = s.DB.Model(&model.School{}).Where("schoolName = ?", name).Count(&cnt).Error
	if cnt == 0 {
		return false, err
	}
	return true, err
}

func (s *School) FindByMailSuffix(suffix string) (sch *model.School, err error) {
	err = s.DB.Model(&model.School{}).Where("mailSuffix = ?", suffix).First(&sch).Error
	return
}

func (s *School) CreateSchool(sch *model.School) (err error) {
	err = s.DB.Model(&model.School{}).Create(&sch).Error
	return
}

func (s *School) UpdateSchool(id int, sch *model.School) (err error) {
	err = s.DB.Model(&model.School{}).Where("schoolID = ?", id).Updates(&sch).Error
	return
}

func (s *School) DeleteSchool(id int) (err error) {
	err = s.DB.Model(&model.School{}).Where("schoolID = ?", id).Delete(&model.School{}).Error
	return
}
