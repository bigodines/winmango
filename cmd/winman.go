package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bigodines/winmango/handler"
	"github.com/julienschmidt/httprouter"
)

var (
	port       = flag.Int("port", 7950, "Winman service port")
	statusPort = flag.Int("statusport", 7960, "Winman status port")
)

func Win(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func main() {
	fmt.Println("Starting service on port: %d", *port)
	statusApi := httprouter.New()

	statusApi.GET("/status", handler.Status.Handle)

	winApi := httprouter.new()

	winApi.GET("/api/v1/win", handler.Win.Handle)

	servers := make([]*http.Server, 0, 3)
	servers = append(servers, router.ServiceApi.Server(addr(*port)))
	servers = append(servers, router.StatusApi.Server(addr(*statusPort)))

	log.Fatal(http.ListenAndServe(":7950", router))
}
