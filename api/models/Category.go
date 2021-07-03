package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type Category struct {
	ID         uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Categories string `gorm:"unique;size:100;not null" json:"categories"`
}

func (c *Category) Prepare() {
	c.ID = 0
	c.Categories = html.EscapeString(strings.TrimSpace(c.Categories))
}

func (c *Category) Validate() error {
	if c.Categories == "" {
		return errors.New("Required Categories")
	}
	return nil
}

func (c *Category) SaveCategory(db *gorm.DB) (*Category, error) {
	var err error
	err = db.Debug().Model(&Category{}).Create(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) DeleteCategory(db *gorm.DB, id string) (int64, error) {
	db = db.Debug().Model(&Category{}).Where("id = ?", id).Take(&Category{}).Delete(&Category{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Category not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Category) FindAllCategories(db *gorm.DB) (*[]Category, error) {
	var err error
	var categories []Category
	err = db.Debug().Model(&Category{}).Limit(100).Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}