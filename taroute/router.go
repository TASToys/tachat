package taroute

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandle)
	r.HandleFunc("/channel/{channel}", ChatHandler)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

type Channel struct {
	Name string
}

type HomeData struct {
	Title    string
	Channels []Channel
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/home.html"))
	homedata := HomeData{
		Title: "blarg",
		Channels: []Channel{
			{Name: "illyohs"},
			{Name: "tinahacks"},
			{Name: "grandpoobear"},
			{Name: "twitchpresents"},
		},
	}

	tmpl.Execute(w, homedata)

}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/chat.html"))

	tmpl.Execute(w, nil)
}
