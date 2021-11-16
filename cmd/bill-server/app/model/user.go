package model

type UserRegisterByEmailServiceArgs struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterByPhoneServiceArgs struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserRegisterResult struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	// Phone    string `json:"phone"`
}

type UserRegisterByEmailApiArgs struct {
	Email    string `v:"required#邮箱不能为空"`
	Password string `v:"required#密码不能为空"`
}

type UserLoginApiArgs struct {
	Email    string `v:"required#邮箱不能为空"`
	Password string `v:"required#密码不能为空"`
}