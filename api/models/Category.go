package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
	"time"
)

type Category struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `grom:"size:255; not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdateAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}

func (cat *Category) Prepare(){
	cat.ID = 0
	cat.Name = html.EscapeString(strings.TrimSpace(cat.Name))
	cat.CreatedAt = time.Now()
	cat.UpdateAt = time.Now()
}

func(cat *Category) Validate() error{
	if cat.Name == ""{
		return errors.New("required category name")
	}
	return nil
}

func (cat *Category) SaveCategory(db *gorm.DB) (*Category, error){
	var err error
	err = db.Debug().Model(&Category{}).Create(&cat).Error
	if err != nil{
		return &Category{}, err
	}
	return cat, nil
}

func(cat *Category) FindAllCategory(db *gorm.DB)(*[]Category, error){
	var err error
	categories := []Category{}
	err = db.Debug().Model(&Category{}).Limit(100).Find(&categories).Error
	if err != nil{
		return &[]Category{}, err
	}
	return &categories, nil
}

func (cat *Category) FindCategoryByID(db *gorm.DB, pid uint64) (*Category, error){
	var err error
	err = db.Debug().Model(&Category{}).Where("id=?",pid).Take(&cat).Error
	if err != nil{
		return &Category{}, err
	}
	return cat, nil
}

func (cat *Category) UpdateCategory(db *gorm.DB) (*Category, error){
	var err error
	err = db.Debug().Model(&Category{}).Where("id=?", cat.ID).Updates(Category{Name:cat.Name,UpdateAt:time.Now()}).Error
	if err != nil{
		return &Category{}, err
	}
	return cat, nil
}

func (cat *Category) DeleteACategory(db *gorm.DB, pid uint64)(int64, error){
	db = db.Debug().Model(&Category{}).Where("id=?",pid).Take(&Category{}).Delete(&Category{})
	if db.Error != nil{
		if gorm.IsRecordNotFoundError(db.Error){
			return 0, errors.New("category not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}