package main

import (
	"os"
	"net/http"
	"strings"
	"io/ioutil"
	"testing"
	"log"
	"fmt"
)

func TestMain(m *testing.M) {
	os.Getpid()
}

func TestHello(t *testing.T) {
	fmt.Println("Hello")
}

func TestString(t *testing.T) {
	addr := ProvideListeningAddr()
	log.Println(addr)

	resp, err := http.Get("http://" + addr + ":4000/string")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	b := strings.Builder { }
	b.Write(body)
	if strings.Compare(b.String(), Always) != 0 {
		t.Fail()
	}
}
