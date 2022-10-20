package models

import "time"

type Address struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User User
	UserID string `gorm:"size:36;Index"`;
	Name        string `gorm:"size:100"`
	IsPrimary   bool
	CityID      string `gorm:"size:100"`
	ProvienceID string `gorm:"size:100"`
	Address1    string `gorm:"size:100"`
	Address2    string `gorm:"size:100"`
	Email       string `gorm:"size:100"`
	PostCode    string `gorm:"size:100"`
	CreatedAt   time.Time
	UpdatedAt time.Time

}