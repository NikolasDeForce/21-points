package main

import (
	"fmt"
	"net/http"
	"os"
	"point/handlers"
	"time"
)

var port = ":8010"

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		port = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         port,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/game", http.HandlerFunc(handlers.GameEngHandler))
	mux.Handle("/game/rus", http.HandlerFunc(handlers.GameRusHandler))
	mux.Handle("/game/check", http.HandlerFunc(handlers.GameParametrsHandler))
	mux.Handle("/game/start/player", http.HandlerFunc(handlers.GamePlayerHandler))
	mux.Handle("/game/start/player/take", http.HandlerFunc(handlers.GameTakeHandler))
	mux.Handle("/game/start/diler", http.HandlerFunc(handlers.GameDilerHandler))
	mux.Handle("/game/final", http.HandlerFunc(handlers.GameFinalHandler))

	fmt.Println("Ready to Serve HTTP to PORT", port)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
