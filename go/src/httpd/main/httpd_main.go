package main

import (
	"os"
	"net/http"
	"log"
	"httpd"
)

/*
	Start the Coursehero Httpd server.
	Usage: httpd [port #] &
*/
func main() {
	addr := httpd.ProvideListeningAddr()
	var port string

	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = httpd.ListeningPort
	}

	log.Println("Listening: " + addr + ":" + port)
	echo := httpd.Echo{ addr }

	http.Handle("/string", httpd.String(httpd.Always))
	http.Handle("/echo", &echo)

	err := http.ListenAndServe(addr + ":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

