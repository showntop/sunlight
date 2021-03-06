package models

import (
	"fmt"

	"github.com/satori/go.uuid"
)

const (
	KEY_NAMESPACE = "sun-token:"
)

type Token struct {
}

func CreateTokenFor(user *User) string {
	token := uuid.NewV4().String()
	StoreM.Cache.Add(fmt.Sprintf("%s%s", KEY_NAMESPACE, token), *user)
	return token
}

func RetriveUserFromToken(token string) (*User, error) {
	value, ok := StoreM.Cache.Get(fmt.Sprintf("%s%s", KEY_NAMESPACE, token))
	if !ok {
		return nil, fmt.Errorf("unrecgonized toke %s", token)
	}
	user := value.(User)
	return &user, nil
}
