package models

import (
	"time"
)

type TbMblog struct {
	ID          uint32    `gorm:"primaryKey;column:id;type:int(10) unsigned;not null" json:"-"`
	BlogID      string    `gorm:"column:blogId;type:varchar(200);not null;default:''" json:"blogId"`
	Name        string    `gorm:"column:name;type:varchar(255);not null;default:''" json:"name"`
	Text        string    `gorm:"column:text;type:varchar(3000);not null;default:''" json:"text"`
	Imgs        string    `gorm:"column:imgs;type:varchar(3000);not null;default:''" json:"imgs"`
	Scheme      string    `gorm:"column:scheme;type:varchar(300);not null;default:''" json:"scheme"`
	TimeCreated time.Time `gorm:"column:time_created;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"timeCreated"`
}