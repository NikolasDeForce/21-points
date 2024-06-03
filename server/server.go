package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	core "point/core"
	"text/template"
	"time"
)

type Content struct {
	Title string
	Text  string
}

type Final struct {
	Winner      core.Winner
	PlayerMoney core.Players
	BotMoney    core.Bot
}

var p = core.Players{
	PlayerName: "Player",
	Money:      2000,
	Bet:        200,
	Card:       "",
	Points:     0,
}

var b = core.Bot{
	DefName: "Diler",
	Money:   2000,
	Bet:     200,
	Card:    "",
	Points:  0,
}

var win = core.Winner{
	WinnerName: "",
}

var resPlayer, resBot, winner = core.StartGame(&p, &b, &win)

var port = ":8010"

func gameEngHandler(w http.ResponseWriter, r *http.Request) {
	resPlayer, resBot, winner = core.StartGame(&p, &b, &win)

	tmpl, _ := template.ParseFiles("gameeng.html")

	content := Content{
		Title: "This is Russian Game '21 Point",
		Text:  "Please, read a role, before press a button",
	}

	err := tmpl.Execute(w, content)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gameRusHandler(w http.ResponseWriter, r *http.Request) {
	resPlayer, resBot, winner = core.StartGame(&p, &b, &win)

	tmpl, _ := template.ParseFiles("gamerus.html")

	content := Content{
		Title: "Игра '21 Очко'",
		Text:  "Перед тем как нажать кнопку, пожалуйста, прочитайте правила игры!",
	}

	err := tmpl.Execute(w, content)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gamePlayerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("gameplayer.html")

	err := tmpl.Execute(w, resPlayer)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gameDilerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("gamediler.html")

	err := tmpl.Execute(w, resBot)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gameFinalHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("gamefinal.html")

	final := Final{
		Winner:      winner,
		PlayerMoney: resPlayer,
		BotMoney:    resBot,
	}

	err := tmpl.Execute(w, final)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

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

	mux.Handle("/game/eng", http.HandlerFunc(gameEngHandler))
	mux.Handle("/game/rus", http.HandlerFunc(gameRusHandler))
	mux.Handle("/game/start/player", http.HandlerFunc(gamePlayerHandler))
	mux.Handle("/game/start/diler", http.HandlerFunc(gameDilerHandler))
	mux.Handle("/game/final", http.HandlerFunc(gameFinalHandler))

	fmt.Println("Ready to Serve HTTP to PORT", port)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
