package models

import "time"

type RepositoryPool struct {
	Id       int    `xorm:"'id' pk"`
	Identity string `xorm:"'identity'"`
	Hash     string `xorm:"'hash'"`
	Name     string `xorm:"'name'"`
	Ext      string `xorm:"'ext'"`
	Size     int    `xorm:"'size'"`
	Path     string `xorm:"'path'"`

	CreatedAt time.Time `xorm:"'created_at' created"`
	UpdatedAt time.Time `xorm:"'updated_at' updated"`
	DeletedAt time.Time `xorm:"'deleted_at' deleted"`
}

var RepositoryPoolName = "repository_pool"

// TableName 会将 UserBasic 的表名重写为 `user_basic`
func (*RepositoryPool) TableName() string {
	return RepositoryPoolName
}

var pool *RepositoryPool

func NewPoolModel() *RepositoryPool {
	if pool == nil {
		pool = &RepositoryPool{}
	}
	return pool
}
