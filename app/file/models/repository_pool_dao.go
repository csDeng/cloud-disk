package models

import "xorm.io/xorm"

type PoolDao interface {
	// 增
	AddFile(Engine *xorm.Engine, p *RepositoryPool) (bool, error)

	// 删

	// 改

	// 查
	// 检查文件是否在存储池
	CheckFileIfExisted(Engine *xorm.Engine, hash string) (*RepositoryPool, error)
}

func (m *RepositoryPool) AddFile(Engine *xorm.Engine, p *RepositoryPool) (bool, error) {
	ok, err := Engine.Insert(p)
	if err != nil {
		return false, err
	}
	return ok > 0, nil
}

func (m *RepositoryPool) CheckFileIfExisted(Engine *xorm.Engine, hash string) (*RepositoryPool, error) {
	rp := new(RepositoryPool)
	_, err := Engine.Where("hash = ?", hash).Get(rp)
	if err != nil {
		return nil, err
	}
	return rp, nil
}
