package models

import "time"

type UserRepository struct {
	Id                 int    `xorm:"id"`
	Identity           string `xorm:"identity"`
	ParentId           int    `xorm:"parent_id"`
	UserIdentity       string `xorm:"user_identity"`
	RepositoryIdentity string `xorm:"repository_identity"`
	Ext                string `xorm:"ext"`
	Name               string `xorm:"name"`

	CreatedAt time.Time `xorm:"'created_at' created"`
	UpdatedAt time.Time `xorm:"'updated_at' updated"`
	DeletedAt time.Time `xorm:"'deleted_at' deleted"`
}

var UserRepositoryName = "user_repository"

func (*UserRepository) TableName() string {
	return UserRepositoryName
}
