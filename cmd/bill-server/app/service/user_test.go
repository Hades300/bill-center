package service

import (
	"bill-server/app/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// 注册用户
	user := model.UserRegisterByEmailServiceArgs{}
	user.Password = "test"
	user.Email = "test@qq.com"
	if err := User.RegisterByEmail(user); err != nil {
		b,_:=json.Marshal(err)
		fmt.Print(string(b))
	}
}
