package dao

type ProjectModel struct {
	Id     int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id" binding:"required"`
	Name   string `gorm:"column:name;type:varchar(30)" json:"name" binding:"required"`
	Intro  string `gorm:"column:intro;type:varchar(100)" json:"intro" binding:"required"`
	Time   string `gorm:"column:time;type:varchar(50)" json:"time" binding:"required"`
	Count  int    `gorm:"column:count;type:int(11)" json:"count" binding:"required"`
	TeamId int    `gorm:"column:team_id;type:int(11)" json:"team_id" binding:"required"`
}
