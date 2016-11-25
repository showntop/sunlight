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
