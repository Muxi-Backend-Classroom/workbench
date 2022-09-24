package dao

import (
	"workbench/model"
)

type TeamModel struct {
	Id           uint32 `gorm:"column:id;not null" binding:"required"`
	Name         string `gorm:"column:name;" binding:"required"`
	Count		 uint32	`gorm:"column:count;not null" binding:"required"`
	Time         string `gorm:"column:time;not null" binding:"required"`
	Creator 	 string `gorm:"column:creator;not null" binding:"required"`
}

type GroupModel struct {
	Id           uint32 `gorm:"column:id;not null" binding:"required"`
	Name         string `gorm:"column:name;" binding:"required"`
	Order		 uint32	`gorm:"column:order;not null" binding:"required"`
	Count        uint32 `gorm:"column:count;not null" binding:"required"`
	Leader 	 	 string `gorm:"column:leader;not null" binding:"required"`
	Time 	 	 string `gorm:"column:time;not null" binding:"required"`
}

type ApplyModel struct {
	Id           uint32 `gorm:"column:id;not null" binding:"required"`
	UserId		 uint32 `gorm:"column:user_id;not null" binding:"required"`
}

type UserModel struct {
	Id           uint32 `gorm:"column:id;not null" binding:"required"`
	Name         string `gorm:"column:name;" binding:"required"`
	RealName     string `gorm:"column:real_name;" binding:"required"`
	Email        string `gorm:"column:email;" binding:"required"`
	Avatar       string `gorm:"column:avatar;" binding:"required"`
	Tel          string `gorm:"column:tel;"`
	Role         uint32 `gorm:"column:role;" binding:"required"`
	TeamId       uint32 `gorm:"column:team_id;" binding:"required"`
	GroupId      uint32 `gorm:"column:group_id;" binding:"required"`
	EmailService uint32 `gorm:"column:email_service;" binding:"required"`
	Message      uint32 `gorm:"column:message;" binding:"required"`
}

func (t *TeamModel) TableName() string {
	return "teams"
}

func (g *GroupModel) TableName() string {
	return "groups"
}

func (a *ApplyModel) TableName() string {
	return "applys"
}

// Create ...
func (t *TeamModel) Create() error {
	return model.DB.Self.Create(t).Error
}

func (g *GroupModel) Create() error {
	return model.DB.Self.Create(g).Error
}

func (a *ApplyModel) Create() error {
	return model.DB.Self.Create(a).Error
}

// Save ...
func (t *TeamModel) Save() error {
	return model.DB.Self.Save(t).Error
}

func (g *GroupModel) Save() error {
	return model.DB.Self.Save(g).Error
}

func (a *ApplyModel) Save() error {
	return model.DB.Self.Save(a).Error
}