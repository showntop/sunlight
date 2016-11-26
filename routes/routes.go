package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/sunlight/handlers"
)

func WrapErrorResp(err *handlers.HttpError) string {
	output := []byte(`{
		"message": "response json error",
		"status": 503
		}`)
	output, _ = json.Marshal(map[string]interface{}{
		"status":  err.Code,
		"message": err.Message,
	})

	return string(output)
}

func Instrument() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/users", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		usersC := new(handlers.Users)
		results, err := usersC.Show(req)
		if err != nil {
			http.Error(rw, WrapErrorResp(err), err.Code)
			// rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.POST("/api/v1/users", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		usersC := new(handlers.Users)
		results, err := usersC.Create(req)
		if err != nil {
			http.Error(rw, WrapErrorResp(err), err.Code)
			// rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	router.PATCH("/api/v1/users", func(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json")
		usersC := new(handlers.Users)
		results, err := usersC.Update(req)
		if err != nil {
			http.Error(rw, WrapErrorResp(err), err.Code)
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
			http.Error(rw, WrapErrorResp(err), err.Code)
			// rw.Write(WrapErrorResp(err))
			return
		}
		rw.Write(results)
	})

	return router
}
