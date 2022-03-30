package main

import (
	"encoding/json"
	"html/template"
	"log"
	yellowRed "myapp/yellowRed"
	"net/http"
	"strconv"
)

func resetYellowRed(w http.ResponseWriter, r *http.Request) {

	yellowRed.ResetGame()
	renderTemplate(w, "yellowRed.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	//i have added here playerChoice variable
	//from object request, we getting variable "c" content
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := yellowRed.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {

	fs := http.FileServer(http.Dir("../Project"))
	http.Handle("/", fs)

	http.HandleFunc("/yellowRed.html", resetYellowRed)
	http.HandleFunc("/yellowRed.html/play", playRound)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
