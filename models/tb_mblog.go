package models

import (
	"time"
)

//TbMBlog
type TbMBlog struct {
	Id          uint       `gorm:"column:id" form:"id" json:"id" comment:"" columnType:"int(10) unsigned" dataType:"int" columnKey:"PRI"`
	BlogId      string     `gorm:"column:blogid" form:"blogid" json:"blogid" comment:"blogid" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
	Name        string     `gorm:"column:name" form:"name" json:"name" comment:"name" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
	Text        string     `gorm:"column:text" form:"text" json:"text" comment:"text" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
	Imgs        string     `gorm:"column:imgs" form:"imgs" json:"imgs" comment:"imgs" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
	Scheme      string     `gorm:"column:scheme" form:"scheme" json:"scheme" comment:"scheme" columnType:"varchar(100)" dataType:"varchar" columnKey:""`
	TimeCreated *time.Time `gorm:"column:time_created" form:"time_created" json:"time_created" comment:"插入时间" columnType:"datetime" dataType:"datetime" columnKey:""`
}

//TableName
func (m *TbMBlog) TableName() string {
	return "tb_mblog"
}

//One
func (m *TbMBlog) One() (one *TbMBlog, err error) {
	one = &TbMBlog{}
	err = crudOne(m, one)
	return
}

//Create
func (m *TbMBlog) Create() (err error) {
	m.Id = 0

	return db.Create(m).Error
}
