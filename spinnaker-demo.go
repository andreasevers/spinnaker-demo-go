package main

import (
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/robfig/cron"
)

type MyHandler struct {
}

// func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Print("Meeseeks")
// 	var path string
// 	path = r.URL.Path[1:]
// 	if path == "" {
// 		path = "index.html"
// 	}
// 	data, err := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "data"}

// 	if err == nil {
// 		w.Write(data)
// 	} else {
// 		w.WriteHeader(404)
// 		w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
// 	}
// }

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "data"})
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
	http.Handle("/", new(MyHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
