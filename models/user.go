package models

import (
	"reflect"
	"strconv"
	"strings"
	// "regexp"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Username       string `json:"username"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Password       string `json:"-"`
	HashedPassword string `json:"-" db:"hashed_password"`

	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description" db:"description"`

	Token string `json:"token"`
}

func (u *User) Validate() error {

	// // [a-z][A-Z][0-9]{8,12}
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

	var fields []string = []string{"username", "hashed_password", "created_at", "updated_at"}
	var vnames []string = []string{":username", ":hashed_password", ":created_at", ":updated_at"}

	if u.Mobile != "" {
		fields = append(fields, "mobile")
		vnames = append(vnames, ":mobile")
	}
	if u.Email != "" {
		fields = append(fields, "email")
		vnames = append(vnames, ":email")
	}

	var id int64
	err := StoreM.Master.QueryRowx(fmt.Sprintf("INSERT INTO users(%s) VALUES(%s) RETURNING id", strings.Join(fields, ","), strings.Join(vnames, ",")), u).Scan(&id)
	if err != nil {
		return err
	}
	u.Id = id
	return nil
}

func GetUserBy(fvalue string) (*User, error) {
	var user *User = &User{}
	var err error
	if strings.Contains(fvalue, "@") { //email
		return nil, nil
	} else { //mobile
		err = StoreM.Master.Get(user, "SELECT users.id, users.username, COALESCE(users.email, '') as email, users.mobile, users.hashed_password, COALESCE(user_profiles.nickname, '') AS nickname, COALESCE(user_profiles.avatar, '') AS avatar, COALESCE(user_profiles.description, '') AS description FROM users LEFT JOIN user_profiles on users.id = user_profiles.user_id WHERE mobile = $1", fvalue)
	}

	return user, err
}

func UpdateUserProfile(userId int64, updateInfo map[string]interface{}) error {
	//allowed update attribute\
	var fields []string
	var values []string
	fields = append(fields, "user_id")
	values = append(values, strconv.FormatInt(userId, 10))
	var updateSubSQL []string
	for field, value := range updateInfo {
		fields = append(fields, field)
		if reflect.TypeOf(value).Kind() == reflect.String {
			updateSubSQL = append(updateSubSQL, field+"='"+value.(string)+"'")
			values = append(values, fmt.Sprintf("'%s'", value.(string)))
		}
		if reflect.TypeOf(value).Kind() == reflect.Int {
			updateSubSQL = append(updateSubSQL, field+"="+strconv.Itoa(value.(int)))
			values = append(values, strconv.Itoa(value.(int)))
		}
	}

	fieldSubSqlStr := strings.Join(fields, ",")
	valueSubSqlStr := strings.Join(values, ",")
	updateSubSQLStr := strings.Join(updateSubSQL, ",")

	sql := fmt.Sprintf(`INSERT INTO user_profiles (%s) VALUES (%s) ON CONFLICT (user_id) DO 
						UPDATE SET %s `, fieldSubSqlStr, valueSubSqlStr, updateSubSQLStr)
	fmt.Println(sql)
	_, err := StoreM.Master.Exec(sql)
	return err
}
