package vars

import (
	"errors"

	grds "github.com/go-redis/redis/v8"
)

var (
	// 邮箱相关
	ErrEmailIsExisted  = errors.New("邮箱已存在")
	ErrEmailNotGetCode = errors.New("当前邮箱没有获取验证码")
	ErrEmailSend       = errors.New("邮件发送失败")

	// 数据库相关
	ErrAdd    = errors.New("新增失败")
	ErrUpdate = errors.New("修改失败")
	ErrDelete = errors.New("删除失败")

	// 用户相关

	ErrLogin   = errors.New("用户名或密码错误")
	ErrRegCode = errors.New("验证码存储失败")

	// token 相关
	ErrRefreshTokenIsNotExisted = errors.New("refresh_token 不存在")
	ErrRefreshTokenHasUsed      = errors.New("refresh_token 已经被使用")
	ErrTokenInvalid             = errors.New("token 无效")

	// redis 相关
	ErrKeyIsNotExisted = grds.Nil

	// lua 相关
	ErrLuaFail = errors.New("lua 执行失败")

	// 文件相关
	ErrFileHasExisted = errors.New("当前文件已处在")
	ErrFileCreate     = errors.New("文件创建失败")
)
