// Code generated by goctl. DO NOT EDIT.
package types

type UserRegisterRequest struct {
	Code     string `json:"code"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type MailRegisterRequest struct {
	Email string `json:"email"`
}

type MailRegisterResponse struct {
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
