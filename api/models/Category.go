package models

type Category struct {
	ID         uint32 `gorm:"primary_key;not null;unique" json:"id"`
	Categories string `gorm:"unique" json:"categories"`
}
