package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID         string `gorm:"size:36:Index"`
	User             User
	UserID           string `gorm:"size:36:Index"`
	ProductImage []ProductImage
	Catagories []Catagory `gorm:"many2many:product_catagories;"`
	Sku              string `gorm:"size:100:Index"`
	Slug             string `gorm:"size:255"`
	Name             string `gorm:"size:255"`
	Price            decimal.Decimal
	Stock            int
	Weight           decimal.Decimal
	ShortDescription string `gorm:"type:text"`
	Description      string `gorm:"type:text"`
	Status           string `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (p *Product) GetProducts(db *gorm.DB, perPage int, page int) (*[]Product, int64, error) {
	var err error
	var product []Product
	var count int64

	err = db.Debug().Model(&Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = db.Debug().Model(&Product{}).Order("created_at desc").Limit(perPage).Offset(offset).Find(&product).Error
	if err != nil {
		return nil, 0, err
	}

	return &product, count, nil
}

func (p *Product) FindBySlug(db *gorm.DB, slug string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Model(&Product{}).Where("slug = ?",slug).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}