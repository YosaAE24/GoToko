package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User        User
	UserID      string `gorm:"size:36;index"`
	Oder        Order
	OderID      string `gorm:"size:36;index"`
	TrackNumber string `gorm:"size:255;index"`
	Status      string `gorm:"size:36;index"`
	TotalQty    int
	TotalWeight decimal.Decimal
	FirstName   string `gorm:"size:100;not null"`
	LastName    string `gorm:"size:100;not null"`
	CityID      string `gorm:"size:100"`
	ProvinceID  string `gorm:"size:100"`
	Address1    string `gorm:"size:100"`
	Address2    string `gorm:"size:100"`
	Phone       string `gorm:"size:50"`
	Email       string `gorm:"size:100"`
	PostCode    string `gorm:"size:100"`
	ShippedBy   string `gorm:"size:36"`
	ShippedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}