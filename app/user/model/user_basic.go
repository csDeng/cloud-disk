package model

import "time"

type UserBasic struct {
	Id       int    `xorm:"'id' pk" `
	Identity string `xorm:"'identity'"`
	Name     string `xorm:"'name'"`
	Password string `xorm:"'password'"`
	Email    string `xorm:"'email'"`

	CreatedAt time.Time `xorm:"'created_at' created"`
	UpdatedAt time.Time `xorm:"'updated_at' updated"`
	DeletedAt time.Time `xorm:"'deleted_at' deleted"`
}

var UserBasicName = "user_basic"

// TableName 会将 UserBasic 的表名重写为 `user_basic`
func (*UserBasic) TableName() string {
	return UserBasicName
}

var userModel *UserBasic

func NewUserModel() *UserBasic {
	if userModel == nil {
		userModel = &UserBasic{}
	}
	return userModel
}
