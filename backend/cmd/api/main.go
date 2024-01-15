package main

import (
	"fmt"
	"log"
	"net/http"

	"todo-backend/cmd/api/router"
	"todo-backend/config"
)

// @title todo-backend api
// @version 1.0
// @description This is a training RESTful API for basic todo app

// @contact.name Pawel Gryc

// @host localhost:8080
// @basePath /v1

func main() {
	c := config.New()
	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
