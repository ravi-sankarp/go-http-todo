package main

import (
	"errors"
	"fmt"
	dbState "http-todo/db"
	"http-todo/handlers"
	"net/http"
)

func main() {
	dbState.ConnectToDb()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.TodoHandler)

	httpHandler := LoggingMiddleware(mux)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", httpHandler)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server shutting down")
	} else if err != nil {
		fmt.Println("Something went wrong", err)
	}

}
