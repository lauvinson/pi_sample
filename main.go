package main

import (
	"log"
	weber "milkomeda.org/milkomeda"
	"milkomeda.org/pi"
	"net/http"
)

func main() {
	p := pi.NewWeb(&weber.Web{})
	p.W.Router(&weber.DefaultRouter{})
	var h = p.W.HandleFunc()
	log.Fatal(http.ListenAndServe(":9999", *h))
}
