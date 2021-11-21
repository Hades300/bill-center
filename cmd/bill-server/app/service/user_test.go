package service

import (
	"encoding/json"
	"fmt"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	// 注册用户
	user := model.UserRegisterByEmailServiceArgs{}
	user.Password = "test"
	user.Email = "test@qq.com"
	if err := User.RegisterByEmail(user); err != nil {
		b, _ := json.Marshal(err)
		fmt.Print(string(b))
	}
}
