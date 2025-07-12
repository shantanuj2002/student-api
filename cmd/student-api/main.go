package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shantanuj2002/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()
	//database setup

	//setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to stident api"))
	})
	//setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Printf("server start %s", cfg.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("fail to start server")
	}

}
