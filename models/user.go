package models

import (
	// "strings"
	// "errors"
	// "regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Username       string `json:"username"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Password       string `json:"password" sql:"-"`
	HashedPassword string `json:"password" sql:"hashed_password"`
}

func (u *User) Validate() error {
	//[a-z][A-Z][0-9]{8,12}
	// if matched, _ := regexp.MatchString(`[a-z][A-Z][0-9]{8,12}`, u.Password); !matched {
	// 	return errors.New("用户名格式错误")
	// }
	// ///((?=.*[a-z])(?=.*\d)|(?=[a-z])(?=.*[#@!~%^&*])|(?=.*\d)(?=.*[#@!~%^&*]))[a-z\d#@!~%^&*]{8,16}/i
	// if matched, _ := regexp.MatchString(`((?=.*[a-z])(?=.*\d)|(?=[a-z])(?=.*[#@!~%^&*])|(?=.*\d)(?=.*[#@!~%^&*]))[a-z\d#@!~%^&*]{8,16}`, u.Password); !matched {
	// 	return errors.New("密码格式错误")
	// }
	return nil
}

func (u *User) EncryptPassword() error {
	// u.HashedPassword
	p, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = string(p)
	return nil
}

func (u *User) Authenticate() error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(u.Password))
}

func CreateUser(u *User) error {
	return nil
}

func GetUserBy(fvalue string) (*User, error) {
	return nil, nil
}
