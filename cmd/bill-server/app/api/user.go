package api

import (
	"bill-server/app/model"
	"bill-server/app/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var User = UserApi{}

type UserApi struct {}

type UserApiI interface {
	Login(r *ghttp.Request)
	Register(r *ghttp.Request)
	Logout(r *ghttp.Request)
}



func (u *UserApi) Login(r *ghttp.Request) {
	var (
		args model.UserLoginApiArgs
		err error
	)
	if err=r.ParseForm(&args);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	var user *model.User
	if user,err=service.User.LoginByEmail(args.Email,args.Password);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	r.Session.Set("user",user)
	JsonSuccessExit(r,"",user)
}



func (u *UserApi) Register(r *ghttp.Request) {
	var (
		args model.UserRegisterByEmailApiArgs
		err error
	)
	if err=r.ParseForm(&args);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	var serviceArgs model.UserRegisterByEmailServiceArgs
	if err:=gconv.Struct(args,&serviceArgs);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	if err=service.User.RegisterByEmail(serviceArgs);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	var user *model.User
	var serviceUser *model.UserRegisterResult
	if user,err=service.User.GetUserByEmail(args.Email);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	if err=gconv.Struct(user,&serviceUser);err!=nil{
		JsonErrExit(r,1,err.Error())
	}
	JsonSuccessExit(r,"",serviceUser)
}