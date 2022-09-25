package model

import (
	"cloud_disk/app/common/vars"
	"fmt"

	"xorm.io/xorm"
)

type UserDao interface {
	// 增
	AddUser(Engine *xorm.Engine, user *UserBasic) (bool, error)

	// 删

	// 改

	// 查

	EmailIfExisted(Engine *xorm.Engine, email string) (bool, error)
	NameIfExisted(Engine *xorm.Engine, name string) (bool, error)
	GetUserByIdentity(Engine *xorm.Engine, identity string) (*UserBasic, error)
	Login(Engine *xorm.Engine, name, password string) (*UserBasic, error)
}

func (m *UserBasic) AddUser(Engine *xorm.Engine, user *UserBasic) (bool, error) {
	u, err := Engine.InsertOne(user)
	if err != nil {
		return false, err
	}
	return u > 0, nil
}

func (m *UserBasic) EmailIfExisted(Engine *xorm.Engine, email string) (bool, error) {
	cnt, err := Engine.Where("email = ?", email).Count(m)
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (m *UserBasic) NameIfExisted(Engine *xorm.Engine, name string) (bool, error) {
	cnt, err := Engine.Where("name = ?", name).Count(m)
	if err != nil {
		return true, err
	}
	return cnt > 0, fmt.Errorf("用户名: %s 已存在", name)

}

func (m *UserBasic) GetUserByIdentity(Engine *xorm.Engine, identity string) (*UserBasic, error) {
	user := new(UserBasic)
	_, err := Engine.Where("identity = ?", identity).Get(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserBasic) Login(Engine *xorm.Engine, name, password string) (*UserBasic, error) {
	user := new(UserBasic)
	has, err := Engine.Where("name = ? AND password = ?", name, password).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, vars.ErrLogin
	}
	return user, nil
}
