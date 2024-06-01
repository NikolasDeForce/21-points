package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	core "point"
	"text/template"
	"time"
)

type Content struct {
	Title string
	Text  string
}

type GamePlayerContent struct {
	NamePlayer   string
	MoneyPlayer  int
	BetPlayer    int
	CardPlayer   string
	PointsPlayer int
	Bot          GameBotContent
	Winner       string
}

type GameBotContent struct {
	NameBot   string
	MoneyBot  int
	BetBot    int
	CardBot   string
	PointsBot int
}

var port = ":8010"

func gameEngHandler(w http.ResponseWriter, r *http.Request) {
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

func gameStartHandler(w http.ResponseWriter, r *http.Request) {
	p := core.Players{
		PlayerName: "Nick",
		Money:      2000,
		Bet:        200,
		Card:       "",
		Points:     0,
	}
	b := core.Bot{
		DefName: "Diler",
		Money:   2000,
		Bet:     200,
		Card:    "",
		Points:  0,
	}
	win := core.Winner{
		WinnerName: "",
	}

	resPlayer, resBot, winner := core.StartGame(&p, &b, &win)

	tmpl, _ := template.ParseFiles("gamestartplayer.html")

	contentPlayer := GamePlayerContent{
		NamePlayer:   resPlayer.PlayerName,
		MoneyPlayer:  resPlayer.Money,
		BetPlayer:    resPlayer.Bet,
		CardPlayer:   resPlayer.Card,
		PointsPlayer: resPlayer.Points,
		Bot: GameBotContent{
			resBot.DefName,
			resBot.Money,
			resBot.Bet,
			resBot.Card,
			resBot.Points,
		},
		Winner: winner.WinnerName,
	}
	err := tmpl.Execute(w, contentPlayer)
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
	mux.Handle("/game/start", http.HandlerFunc(gameStartHandler))

	err := s.ListenAndServe()
	fmt.Println("Ready to Serve HTTP to PORT", port)
	if err != nil {
		fmt.Println(err)
		return
	}
}
