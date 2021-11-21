package service

import (
	"errors"

	"github.com/hades300/bill-center/cmd/bill-server/app/dao"
	"github.com/hades300/bill-center/cmd/bill-server/app/model"
)

var (
	ErrNickNameExist = errors.New("nickname exist")
	ErrEmailExist    = errors.New("email exist")
	ErrUserNotExist  = errors.New("user not exist")
	ErrPhoneExist    = errors.New("phone exist")
	ErrPasswordWrong = errors.New("password wrong")
)

var (
	User = NewUserService()
)

type UserServiceI interface {
	GetUser(id int) (*model.User, error)
	GetUserByNickName(name string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByPhone(phone string) (*model.User, error)

	isNameExists(name string) bool
	isEmailExists(email string) bool
	isPhoneExists(phone string) bool

	LoginByEmail(email, password string) (*model.User, error)
	RegisterByEmail(args model.UserRegisterByEmailServiceArgs) error
	RegisterByPhone(args model.UserRegisterByPhoneServiceArgs) error
}
type userService struct{}

var _ UserServiceI = &userService{}

func NewUserService() UserServiceI {
	return &userService{}
}

func (u *userService) LoginByEmail(email, password string) (*model.User, error) {
	user, err := u.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, ErrPasswordWrong
	}
	return user, nil
}

func (u *userService) RegisterByEmail(args model.UserRegisterByEmailServiceArgs) error {
	if u.isEmailExists(args.Email) {
		return ErrEmailExist
	}
	_, err := dao.User.Ctx(nil).Insert(args)
	return err
}

func (u *userService) RegisterByPhone(args model.UserRegisterByPhoneServiceArgs) error {
	if u.isPhoneExists(args.Phone) {
		return ErrPhoneExist
	}
	_, err := dao.User.Ctx(nil).Insert(args)
	return err
}

func (u *userService) GetUser(id int) (*model.User, error) {
	var user model.User
	err := dao.User.Ctx(nil).Where("id=?", id).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) GetUserByNickName(name string) (*model.User, error) {
	var user model.User
	err := dao.User.Ctx(nil).Where("nickname=?", name).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := dao.User.Ctx(nil).Where("email=?", email).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) GetUserByPhone(phone string) (*model.User, error) {
	var user model.User
	err := dao.User.Ctx(nil).Where("phone=?", phone).Scan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userService) isNameExists(name string) bool {
	cnt, err := dao.User.Ctx(nil).Count("nickname=?", name)
	if err != nil {
		return false
	}
	return cnt != 0
}

func (u *userService) isEmailExists(email string) bool {
	cnt, err := dao.User.Ctx(nil).Count("email=?", email)
	if err != nil {
		return false
	}
	return cnt != 0
}

func (u *userService) isPhoneExists(phone string) bool {
	cnt, err := dao.User.Ctx(nil).Count("phone=?", phone)
	if err != nil {
		return false
	}
	return cnt != 0
}
