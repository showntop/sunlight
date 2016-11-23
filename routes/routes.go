package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/sunlight/handlers"
)

func WrapErrorResp(err *handlers.HttpError) []byte {
	output := []byte(`{
		"message": "response json error",
		"state_code": 503
		}`)
	output, _ = json.Marshal(map[string]interface{}{
		"state_code": err.Code,
		"message":    err.Message,
	})

	return output
}

func Instrument() *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/users", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		usersC := new(handlers.Users)
		results, err := usersC.Create(req)
		if err != nil {
			http.Error(rw, err.Error(), err.Code)
			// rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.POST("/api/v1/sessions", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		sessionsC := new(handlers.Sessions)
		results, err := sessionsC.Create(req)
		if err != nil {
			http.Error(rw, err.Error(), err.Code)
			// rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	return router
}
