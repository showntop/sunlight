package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/showntop/sunlight/models"
)

type Users struct {
	application
}

func (u *Users) Show(req *http.Request) ([]byte, *HttpError) {
	user, err := models.RetriveUserFromToken(req.Header.Get("Sun-Token"))
	if err != nil {
		return nil, AuthErr
	}

	user.Token = req.Header.Get("Sun-Token")
	output, err := json.Marshal(WrapRespData(user))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}

func (u *Users) Create(req *http.Request) ([]byte, *HttpError) {
	//request do
	signupInfo := &struct {
		Username string `json:"username,omitempty"`
		Email    string `json:"email,omitempty"`
		Mobile   string `json:"mobile,omitempty"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&signupInfo)
	if err != nil {
		return nil, BadRequestErr
	}
	//only allowed name field
	//did has the better way or separate into reqmodel  sqlmodel  repmodel
	user := &models.User{
		Username: signupInfo.Username,
		Password: signupInfo.Password,
		Mobile:   signupInfo.Mobile,
		Email:    signupInfo.Email,
	}
	if err := user.EncryptPassword(); err != nil {
		return nil, ServerErr
	}
	if verr := user.Validate(); verr != nil {
		return nil, &HttpError{Code: 402, Message: verr.Error()}
	}

	//save
	err = models.CreateUser(user)
	if err != nil {
		log.Error("server error:", err)
		return nil, DBErr
	}
	user.Token = models.CreateTokenFor(user)
	//respose do
	output, err := json.Marshal(WrapRespData(user))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}

func (u *Users) Update(req *http.Request) ([]byte, *HttpError) {
	user, err := u.AuthUser(req)
	if err != nil {
		return nil, IncorrectAccountErr
	}
	u.CurrentUser = user
	var updateInfo map[string]interface{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&updateInfo)
	if err != nil {
		return nil, BadRequestErr
	}

	err = models.UpdateUserProfile(u.CurrentUser.Id, updateInfo)
	if err != nil {
		log.Error("server error:", err)
		return nil, DBErr
	}
	output, err := json.Marshal(WrapRespData(map[string]interface{}{
		"message": "update success",
	}))
	if err != nil {
		return output, BadRespErr
	}
	return output, nil
}
