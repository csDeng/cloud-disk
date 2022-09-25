package models

import "time"

type ShareBasic struct {
	Id                     int    `xorm:"id"`
	Identity               string `xorm:"identity"`
	UserIdentity           string `xorm:"user_identity"`
	RepositoryIdentity     string `xorm:"repository_identity"`
	UserRepositoryIdentity string `xorm:"user_repository_identity"`
	ExpiredTime            int    `xorm:"expired_time"`
	ClickNum               int    `xorm:"click_num"`

	CreatedAt time.Time `xorm:"'created_at' created"`
	UpdatedAt time.Time `xorm:"'updated_at' updated"`
	DeletedAt time.Time `xorm:"'deleted_at' deleted"`
}

var ShareBasicName = "share_basic"

func (*ShareBasic) TableName() string {
	return ShareBasicName
}
