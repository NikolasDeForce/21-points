package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	core "point/core"
	"strconv"
	"text/template"
	"time"
)

type Content struct {
	Title string
	Text  string
}

type Final struct {
	Winner core.Winner
	Player core.Players
	Bot    core.Bot
}

var p = core.Players{
	PlayerName: "",
	Money:      2000,
	Bet:        0,
	Card:       "",
	Points:     0,
}

var b = core.Bot{
	DefName: "Diler",
	Money:   2000,
	Bet:     0,
	Card:    "",
	Points:  0,
}

var win = core.Winner{
	WinnerName: "",
}

var resPlayer, resBot, winner = core.StartGame(&p, &b, &win)

var port = ":8010"

func gameEngHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("gameeng.html")

	content := Content{
		Title: "This is Russian Game '21 Point",
		Text:  "Please, read a rule, before press a button",
	}

	err := tmpl.Execute(w, content)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gameRusHandler(w http.ResponseWriter, r *http.Request) {
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

func gameParametrsHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	p.PlayerName = r.FormValue("userName")
	if p.PlayerName == "" {
		p.PlayerName = "Player"
	}
	p.Bet, _ = strconv.Atoi(r.FormValue("userBet"))
	if p.Bet == 0 {
		p.Bet = 200
	}
	if p.Bet > 2000 {
		p.Bet = 200
	}

	b.Bet, _ = strconv.Atoi(r.FormValue("userBet"))
	if b.Bet == 0 {
		b.Bet = 200
	}
	if b.Bet > 2000 {
		b.Bet = 200
	}

	if winner.WinnerName == "" {
		winner.WinnerName = resPlayer.PlayerName
	}

	resPlayer, resBot, winner = core.StartGame(&p, &b, &win)

	tmpl, _ := template.ParseFiles("gamecheck.html")

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gamePlayerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("gameplayer.html")

	r.ParseForm()

	err := tmpl.Execute(w, resPlayer)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Hello from %v", r.Host)
	w.WriteHeader(http.StatusOK)
}

func gameTakeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("gametake.html")

	resPlayer = core.TakeOne(&resPlayer)

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
		Winner: winner,
		Player: resPlayer,
		Bot:    resBot,
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

	mux.Handle("/game", http.HandlerFunc(gameEngHandler))
	mux.Handle("/game/rus", http.HandlerFunc(gameRusHandler))
	mux.Handle("/game/check", http.HandlerFunc(gameParametrsHandler))
	mux.Handle("/game/start/player", http.HandlerFunc(gamePlayerHandler))
	mux.Handle("/game/start/player/take", http.HandlerFunc(gameTakeHandler))
	mux.Handle("/game/start/diler", http.HandlerFunc(gameDilerHandler))
	mux.Handle("/game/final", http.HandlerFunc(gameFinalHandler))

	fmt.Println("Ready to Serve HTTP to PORT", port)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
