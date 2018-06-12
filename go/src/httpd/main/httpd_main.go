package main

import (
	"httpd"
	"net/http"
	"log"
)

func main() {
	addr := httpd.ProvideListeningAddr()
	log.Println("Listening: " + addr)
	echo := httpd.Echo{ addr }

	http.Handle("/string", httpd.String(httpd.Always))
	http.Handle("/echo", &echo)

	err := http.ListenAndServe(addr + ":" + httpd.ListeningPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

