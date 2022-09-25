package models

import (
	"cloud_disk/app/common/vars"

	"xorm.io/xorm"
)

type UserRepositoryDao interface {
	// 增
	// 增加用户关联文件
	AddFile(Engine *xorm.Engine, userIdentity string, file *RepositoryPool) (*UserRepository, error)

	// 删

	// 改

	// 查
	// 检查当前用户是否有当前文件
	CheckUserIfHasFile(Engine *xorm.Engine, userIdentity, fileIdentity string) (*UserRepository, error)
}

func (m *UserRepository) AddFile(Engine *xorm.Engine, up *UserRepository) (*UserRepository, error) {
	ok, err := Engine.Insert(up)
	if err != nil {
		return nil, err
	}
	if ok <= 0 {
		return nil, vars.ErrAdd
	}
	return up, nil

}

func (m *UserRepository) CheckUserIfHasFile(Engine *xorm.Engine, userIdentity, fileIdentity string) (*UserRepository, error) {
	up := new(UserRepository)
	_, err := Engine.Where("user_identity = ?", userIdentity).And("repository_identity = ?", fileIdentity).Get(up)
	if err != nil {
		return nil, err
	}
	return up, nil
}
