package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bigodines/winmango/handler"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/julienschmidt/httprouter"
)

var (
	port       = flag.Int("port", 7950, "Winman service port")
	statusPort = flag.Int("statusport", 7960, "Winman status port")
)

func main() {
	fmt.Println("Starting service on port: %d", *port)
	// Setup /status
	statusApi := httprouter.New()
	statusApi.GET("/status", handler.Status)

	// Setup /api/v1
	winApi := httprouter.New()
	winApi.GET("/api/v1/win", handler.Win)

	servers := make([]*http.Server, 0, 3)
	servers = append(servers, &http.Server{Addr: addr(*statusPort), Handler: statusApi})
	servers = append(servers, &http.Server{Addr: addr(*port), Handler: winApi})

	if err := gracehttp.Serve(servers...); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}

// addr returns the string representation of the service address for the http.Server instance.
func addr(p int) string {
	return ":" + strconv.Itoa(p)
}
