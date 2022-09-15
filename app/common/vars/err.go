package vars

import "errors"

var (
	EmailIsExistedErr = errors.New("邮箱已存在")
	AddErr            = errors.New("新增失败")
	LoginErr          = errors.New("用户名或密码错误")
)
