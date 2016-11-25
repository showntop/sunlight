package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/showntop/sunlight/models"
)

type Sessions struct {
}

func (s *Sessions) Create(req *http.Request) ([]byte, *HttpError) {
	loginInfo := &struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&loginInfo)
	if err != nil {
		log.Warnln("session parse request,", err)
		return nil, BadRequestErr
	}

	user, err := models.GetUserBy(loginInfo.Mobile)
	if err != nil {
		log.Errorln("session query database,", err)
		return nil, DBErr
	}
	if user == nil {
		return nil, IncorrectAccountErr
	}
	user.Password = loginInfo.Password
	if err := user.Authenticate(); err != nil {
		log.Warnln("session auth user,", err)
		return nil, IncorrectAccountErr
	}

	//respose do
	output, err := json.Marshal(WrapRespData(user))
	if err != nil {
		log.Warnln("session marshal user,", err)
		return output, BadRespErr
	}
	return output, nil
}
