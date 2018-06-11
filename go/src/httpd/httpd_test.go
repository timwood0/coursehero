package httpd

import (
	"os"
	"os/exec"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
	"testing"
	"log"
)

var addr string = ProvideListeningAddr()
var errorMsg string = "Error 400"

func TestMain(m *testing.M) {
	log.Println(addr)
	// Start up the server
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("GOPATH not set in environment.")
	}

	httpdCmd := exec.Command(gopath + "/bin/httpd")
	err := httpdCmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Allow startup and run the tests
	// XXX Uses the server already running, if any
	// XXX Race-y approach, would rather poll the socket
	exec.Command("sleep", "1").Run()
	runStatus := m.Run()

	// Stop the server
	httpdCmd.Process.Kill()
	if err != nil {
		log.Println(err) // Warning
	}

	os.Exit(runStatus)
}

func getResponseText(resp *http.Response, t *testing.T) string {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	b := strings.Builder { }
	b.Write(body)
	return b.String()
}

func TestString(t *testing.T) {
	resp, err := http.Get("http://" + addr + ":4000/string")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText := getResponseText(resp, t)
	if strings.Compare(bodyText, Always) != 0 {
		t.Fatal(bodyText)
	}
}

func TestStringParams(t *testing.T) {
	resp, err := http.Get("http://" + addr + ":4000/string?param1=what&param2=who")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText := getResponseText(resp, t)
	if strings.Compare(bodyText, Always) != 0 {
		t.Fatal(bodyText)
	}
}

func TestEcho(t *testing.T) {
	testMessage := "When in the course of human events"
	resp, err := http.Get("http://" + addr + ":4000/echo?message=" + url.QueryEscape(testMessage))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText := getResponseText(resp, t)
	if strings.Compare(bodyText, addr + " " + testMessage) != 0 {
		t.Fatal(bodyText)
	}
}

func TestEchoNoParam(t *testing.T) {
	resp, err := http.Get("http://" + addr + ":4000/echo")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText := getResponseText(resp, t)
	if strings.Compare(bodyText, errorMsg) != 0 {
		t.Fatal(bodyText)
	}
}

func TestEchoBadParam(t *testing.T) {
	resp, err := http.Get("http://" + addr + ":4000/echo?var=" + url.QueryEscape("An inconvenient message"))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText := getResponseText(resp, t)
	if strings.Compare(bodyText, errorMsg) != 0 {
		t.Fatal(bodyText)
	}
}

func TestEchoExtraParam(t *testing.T) {
	testMessage := "When in the course of human events"
	resp, err := http.Get("http://" + addr + ":4000/echo?other=Something&message=" + url.QueryEscape(testMessage))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText := getResponseText(resp, t)
	if strings.Compare(bodyText, addr + " " + testMessage) != 0 {
		t.Fatal(bodyText)
	}
}

