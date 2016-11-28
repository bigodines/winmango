package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Win(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	res.Header().Add("Content-Type", "application/json")
	res.Write([]byte(`{"Status": "OK"}`))
}
