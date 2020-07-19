package models

import "github.com/jinzhu/gorm"

type SampleProduct struct {
	gorm.Model
	Name		string	`gorm:"type:varchar(100)"`
	Code		string	`gorm:"type:varchar(100)"`
	ImageUrl	string	`gorm:"type:varchar(100)"`
	Description string	`gorm:"type:text"`
}
type SampleGroup struct {
	gorm.Model
	Name			string			`gorm:"type:varchar(100)"`
	Code			string			`gorm:"type:varchar(100)"`
	SampleProductID	uint			`gorm:"type:int"`
	SampleProducts	[]SampleProduct	`gorm:"foreignkey:ID;association_foreignkey:SampleProductID"`
}