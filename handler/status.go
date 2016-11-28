package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Status(res http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	res.Write([]byte(`{"Status": "OK"}`))
}
