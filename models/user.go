package models

import (
	"strings"
	// "errors"
	// "regexp"
	"database/sql"
	// "fmt"
	"time"

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
	// fmt.Printf("\n %s-%s \n", u.HashedPassword, u.Password)
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(u.Password))
}

func CreateUser(u *User) error {
	u.CreatedAt, u.UpdatedAt = time.Now(), time.Now()
	u.Username = "u" + u.Mobile
	var id int64
	err := StoreM.Master.QueryRow("INSERT INTO users(email,mobile,username,hashed_password,created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id", u.Email, u.Mobile, u.Username, u.HashedPassword, u.CreatedAt, u.UpdatedAt).Scan(&id)
	if err != nil {
		return err
	}
	u.Id = id
	return nil
}

func GetUserBy(fvalue string) (*User, error) {
	var err error
	var id int64
	var username, name, email, mobile, hashedPassword sql.NullString
	if strings.Contains(fvalue, "@") { //email
		return nil, nil
	} else { //mobile
		err = StoreM.Master.QueryRow("SELECT id, username, name, email, mobile, hashed_password FROM users WHERE mobile = $1", fvalue).Scan(&id, &username, &name, &email, &mobile, &hashedPassword)
	}
	user := &User{
		Username:       username.String,
		Name:           name.String,
		Mobile:         mobile.String,
		Email:          email.String,
		HashedPassword: hashedPassword.String,
	}
	user.Id = id
	return user, err
}
