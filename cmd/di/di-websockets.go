package di

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	ws_server "github.com/AntonioMartinezFernandez/golang-websockets-server/pkg/websockets-server"
)

func InitWebsocketsServer() {
	var websocketsAddr = flag.String("addr", ":8080", "http service address")
	flag.Parse()

	hub := ws_server.NewHub()
	go hub.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi from WS server")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws_server.ServeWs(hub, w, r)
	})

	server := &http.Server{
		Addr:              *websocketsAddr,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
