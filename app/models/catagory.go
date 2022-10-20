package models

import "time"

type Catagory struct {
	ID        string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID  string `gorm:"size:36;"`
	Section   Section
	SectionID string  `gorm:"size:36;Index"`
	Products  Product `gorm:"many2many:product_catagories;"`
	Name      string  `gorm:"size:100;"`
	Slug      string  `gorm:"size:100;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}