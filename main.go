package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/HumbertoM10/pokemonapi/packages/handler"
	"github.com/gorilla/mux"
)

// main is in charge of setting the port and executing the server's goroutine.
func main() {
	port := 3000
	router := mux.NewRouter()

	router.HandleFunc("/api/advantage", handler.Advantage).Methods("GET")
	router.HandleFunc("/api/commonMoves", handler.CommonMoves).Methods("GET")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	defer srv.Close()

	// The server is ran using a gorutine so it doesn't block
	go func() {
		log.Printf("Listening on port %d...\n", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Print(err)
		}
	}()

	// Block execution until a sigint signal is recieved
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
