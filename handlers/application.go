package handlers

import (
	"fmt"
	"net/http"

	"github.com/showntop/sunlight/models"
)

type application struct {
	CurrentUser *models.User
}

type HttpError struct {
	Code    int
	Message string
}

func (h *HttpError) Error() string {
	return h.Message
}

var (
	BadRequestErr       = &HttpError{400, "json format error"}
	BadRequestErr2      = &HttpError{422, "invalid query key/value pair"}
	IncorrectAccountErr = &HttpError{403, "用户名或者密码错误"}
	AuthErr             = &HttpError{403, "用户验证失败"}

	ServerErr  = &HttpError{500, "server error"}
	DBErr      = &HttpError{503, "db error"}
	BadRespErr = &HttpError{503, "response format error"}
)

func WrapRespData(data interface{}) interface{} {
	// result := make(map[string]interface{})
	// result["data"] = data
	// result["state_code"] = 200
	// result["message"] = "成功"
	return data
}

func (a *application) AuthUser(req *http.Request) (*models.User, error) {
	token := req.Header.Get("Sun-Token")

	value, ok := models.StoreM.Cache.Get(fmt.Sprintf("%s%s", models.KEY_NAMESPACE, token))
	if !ok {
		return nil, fmt.Errorf("token can not auth %s", "error")
	}
	user := value.(models.User)
	return &user, nil
}
