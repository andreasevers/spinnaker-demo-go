package main

import (
	"log"
	"net/http"

	"github.com/robfig/cron"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//log.Print("Meeseeks")
	var path string
	path = r.URL.Path[1:]
	if path == "" {
		path = "index.html"
	}
	path = "assets/" + string(path)
	data, err := Asset(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}

func logRickAndMortyNames() {
	log.Print("Rick")
	log.Print("Morty")
	log.Print("Beth")
	log.Print("Jerry")
	log.Print("Summer")
	log.Print("Mr. Poopybutthole")
	log.Print("Mr. Meeseeks")
	log.Print("Birdperson")
	log.Print("Ants-In-My-Eyes Johnson")
	log.Print("Jessica")
	log.Print("Gazorpazorpfield")
	log.Print("Noob Noob")
	log.Print("Scary Terry")
	log.Print("Squanchy")
}

func main() {
	c := cron.New()
	c.AddFunc("@every 20s", logRickAndMortyNames)
	c.Start()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
