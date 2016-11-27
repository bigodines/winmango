package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Status(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "{ \"Status\": \"OK\" }")
}

func Win(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func main() {
	fmt.Println("hi")
	router := httprouter.New()
	router.GET("/status", Status)
	router.GET("/api/v1/win", Win)

	log.Fatal(http.ListenAndServe(":7950", router))
}
