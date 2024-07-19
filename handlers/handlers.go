package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"point/core"
	"strconv"
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

var p = core.Players{}

var b = core.Bot{}

var win = core.Winner{}

var resPlayer, resBot = core.StartGame(&p, &b)

func GameEngHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameEngHandler Serving:", r.URL.Path, "from", r.Body)
	tmpl, _ := template.ParseFiles("./templates/gameeng.html")

	content := Content{
		Title: "This is Russian Game '21 Point",
		Text:  "Please, read a rule, before press a button",
	}

	err := tmpl.Execute(w, content)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GameRusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameRusHandler Serving:", r.URL.Path, "from", r.Body)
	tmpl, _ := template.ParseFiles("./templates/gamerus.html")

	content := Content{
		Title: "Игра '21 Очко'",
		Text:  "Перед тем как нажать кнопку, пожалуйста, прочитайте правила игры!",
	}

	err := tmpl.Execute(w, content)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GameParametrsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameParametersHandler Serving:", r.URL.Path, "from", r.Body)

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

	if win.WinnerName == "" {
		win.WinnerName = resPlayer.PlayerName
	}

	resPlayer, resBot = core.StartGame(&p, &b)

	tmpl, _ := template.ParseFiles("./templates/gamecheck.html")

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GamePlayerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GamePlayerHandler Serving:", r.URL.Path, "from", r.Body)
	tmpl, _ := template.ParseFiles("./templates/gameplayer.html")

	r.ParseForm()

	err := tmpl.Execute(w, resPlayer)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GameTakeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameTakeHandler Serving:", r.URL.Path, "from", r.Body)
	tmpl, _ := template.ParseFiles("./templates/gametake.html")

	resPlayer = core.TakeOne(&resPlayer)

	err := tmpl.Execute(w, resPlayer)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GameDilerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameDilerHandler Serving:", r.URL.Path, "from", r.Body)
	tmpl, _ := template.ParseFiles("./templates/gamediler.html")

	err := tmpl.Execute(w, resBot)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GameFinalHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GameFinalHandler Serving:", r.URL.Path, "from", r.Body)
	finalPlayer, finalBot, finalWinner := core.ResultGame(&resPlayer, &resBot, &win)

	tmpl, _ := template.ParseFiles("./templates/gamefinal.html")

	final := Final{
		Winner: finalWinner,
		Player: finalPlayer,
		Bot:    finalBot,
	}

	err := tmpl.Execute(w, final)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}
