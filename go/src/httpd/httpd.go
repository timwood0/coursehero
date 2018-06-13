package httpd

/*
	This package implements a trivial HTTP server.
	It recognizes two hardcoded URL paths:
	/string - Returns a plain, fixed string to the browser; any
		supplied parameters are ignored.
	/echo?message=... - Returns the 'message' parameter value to
		the browser as a plain string.
	This server returns error 400 for URLs that do not conform to
	the above.
*/

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type String string

type Echo struct {
	Tag string
}

var Always string = "This URL always returns the same thing."

// Default; override on command line
var ListeningPort string = "4000"

func rejectRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
	fmt.Fprint(w, "Error 400")
	log.Println("Error 400")
}

// Echo the contents of the 'message' parameter back to the user
// If it does not exist, return a 400 error
func (echo Echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Find the parameters
	err := r.ParseForm()

	if err != nil || r.Form["message"] == nil {
		if err != nil {
			log.Println(err)
		}

		rejectRequest(w)
		return
	}

	// Tag user output with our tag field
	fmt.Fprint(w, echo.Tag + " " + r.Form["message"][0])
}

// Return a constant string to the client.
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func ProvideListeningAddr() string {
	addrs, err := net.InterfaceAddrs()

	// Need at least two interfaces,
	// localhost and 1 other
	// XXX assumes localhost comes first
	if err != nil {
		log.Fatal(err)
	} else {
		if len(addrs) < 2 {
			log.Fatal("No interfaces off machine")
		}
	}

	// We will tag user output with our IP address
	addr := addrs[1].String()
	addr = addr[0:strings.IndexByte(addr, '/')]

	return addr
}

