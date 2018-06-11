package main

import (
	"httpd"
	"net/http"
	"log"
)

// XXX Fixed; prefer configuration
var listeningPort string = "4000"

func main() {
	addr := httpd.ProvideListeningAddr()
	log.Println("Listening: " + addr)
	echo := httpd.Echo{ addr }

	http.Handle("/string", httpd.String(httpd.Always))
	http.Handle("/echo", &echo)

	err := http.ListenAndServe(addr + ":" + listeningPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

