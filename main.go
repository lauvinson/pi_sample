package main

import (
	"log"
	"milkomeda.org/milkomeda"
	pi2 "milkomeda.org/pi"
	"net/http"
)

func main() {
	pi := pi2.New()
	pi.Router(&milkomeda.DefaultRouter{})
	var h = pi.W.HandleFunc()
	log.Fatal(http.ListenAndServe(":9999", *h))
}
