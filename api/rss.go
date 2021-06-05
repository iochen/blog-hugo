package handler

import (
	"net/http"

	"github.com/iochen/trss"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	rss, err := trss.GenerateRSS(name)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("cannot generate rss\n"))
		w.Write([]byte(err.Error()))
		return
	}
	s, err := rss.ToRss()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("cannot generate rss\n"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type","application/xml")
	w.Write([]byte(s))
}
