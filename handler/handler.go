package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Handle(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}
