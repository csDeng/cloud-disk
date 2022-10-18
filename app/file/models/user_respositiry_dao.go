package models

import (
	"cloud_disk/app/common/vars"
	"log"

	"xorm.io/xorm"
)

type HelperType struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int    `json:"size"`
	Path               string `json:"path"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
}

type UserRepositoryDao interface {
	// 增
	// 增加用户关联文件
	AddFile(Engine *xorm.Engine, userIdentity string, file *RepositoryPool) (*UserRepository, error)

	// 创建目录
	CreateFolder(Engine *xorm.Engine, up *UserRepository) (bool, error)

	// 删
	// 用户文件删除
	DelUserFile(Engine *xorm.Engine, identity, userIdentity string) (bool, error)
	// 改
	// 修改文件名称
	EditFileName(Engine *xorm.Engine, userIdentity, identity string) (bool, error)
	// 查
	// 检查当前用户是否有当前文件
	CheckUserIfHasFile(Engine *xorm.Engine, userIdentity, fileIdentity string) (*UserRepository, error)

	GetTotalWithParentId(Engine *xorm.Engine, userIdentity string, parentId int) (int, error)

	GetUserFileList(Engine *xorm.Engine, userIdentity string, parentId, size, offset int) ([]HelperType, error)

	// 检查当前标识的文件是否存在
	CheckIfExistedIdentity(Engine *xorm.Engine, userIdentity, identity string) (*UserRepository, error)

	// 检查当前目录是否存在该名称的文件
	CheckIfHasNameWithPId(Engine *xorm.Engine, parent_id int, user_identity, name string) (bool, error)

	// 根据identity 获取 parent_id
	GetPIdWithIdentity(Engine *xorm.Engine, identity string) (int, error)

	// 检查用户目录文件夹是否存在
	CheckParentIfExisted(Engine *xorm.Engine, userIdentity string, parentId int) (bool, error)

	// 检查当前文件是否是一个目录
	CheckFileIsFolder(Engine *xorm.Engine, userIdentity, identity string) (bool, error)
}

func (m *UserRepository) CheckFileIsFolder(Engine *xorm.Engine, userIdentity, identity string) (bool, error) {
	cnt, err := Engine.Where(
		"identity = ? AND user_identity = ?", identity, userIdentity,
	).And("repository_identity = NULL OR LENGTH(trim(repository_identity)) < 1").Count(m)
	return cnt > 0, err
}

func (m *UserRepository) DelUserFile(Engine *xorm.Engine, identity, userIdentity string) (bool, error) {
	cnt, err := Engine.Where("identity = ? AND user_identity = ?", identity, userIdentity).Delete(m)
	return cnt > 0, err
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

func (m *UserRepository) GetTotalWithParentId(Engine *xorm.Engine, userIdentity string, parentId int) (int, error) {
	cnt, err := Engine.Where("parent_id = ? AND user_identity = ?", parentId, userIdentity).
		Count(m)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (m *UserRepository) GetUserFileList(Engine *xorm.Engine, userIdentity string, parentId, size, offset int) ([]HelperType, error) {
	urName, pName := UserRepositoryName, RepositoryPoolName

	uf := make([]HelperType, 0)
	err := Engine.Table(urName).
		Where("parent_id = ? AND user_identity = ?", parentId, userIdentity).
		Select("user_repository.id, user_repository.name, user_repository.ext, repository_pool.size,repository_pool.path,user_repository.identity, repository_pool.identity AS repository_identity").
		Join("LEFT", pName, "user_repository.repository_identity = repository_pool.Identity").
		Where("user_repository.deleted_at IS NULL").
		Limit(size, offset).Find(&uf)
	if err != nil {
		return nil, err
	}
	return uf, nil
}

func (m *UserRepository) EditFileName(Engine *xorm.Engine, userIdentity, identity, name string) (bool, error) {
	update, err := Engine.Table(m.TableName()).Where("identity = ?", identity).And("user_identity = ?", userIdentity).Update(map[string]interface{}{"name": name})
	if err != nil {
		return false, err
	}
	return update > 0, nil
}

func (m *UserRepository) CheckIfHasNameWithPId(Engine *xorm.Engine, parent_id int, user_identity, name string) (bool, error) {
	cnt, err := Engine.Where("user_identity = ? AND parent_id = ? AND name = ?", user_identity, parent_id, name).Count(m)
	log.Println("del->", cnt, err)
	return cnt > 0, err

}

func (m *UserRepository) CheckIfExistedIdentity(Engine *xorm.Engine, userIdentity, identity string) (*UserRepository, error) {
	user := new(UserRepository)
	_, err := Engine.Where("identity = ?", identity).And("user_identity = ?", userIdentity).Get(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserRepository) GetPIdWithIdentity(Engine *xorm.Engine, identity string) (int, error) {
	user := new(UserRepository)
	_, err := Engine.Where("identity = ?", identity).Get(user)
	if err != nil {
		return 0, err
	}
	return user.ParentId, nil
}

func (m *UserRepository) CheckParentIfExisted(Engine *xorm.Engine, userIdentity string, parentId int) (bool, error) {
	if parentId == 0 {
		// 根目录
		return true, nil
	}
	cnt, err := Engine.Where("id = ?", parentId).And("user_identity = ?", userIdentity).And("repository_identity = NULL").Count(m)
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (m *UserRepository) CreateFolder(Engine *xorm.Engine, up *UserRepository) (bool, error) {
	ok, err := Engine.InsertOne(up)
	if err != nil {
		return false, err
	}
	return ok > 0, nil
}
