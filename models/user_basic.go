package models

type UserBasic struct {
	Id       int
	Name     string
	Password string
	Email    string
}

// TableName 会将 UserBasic 的表名重写为 `user_basic`
func (UserBasic) TableName() string {
	return "user_basic"
}
