package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func quoteGrab(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://bash.org/?random1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	page := html.UnescapeString(string(b))
	if strings.Contains(page, "MySQL daemon") {
		w.WriteHeader(http.StatusGatewayTimeout)
		_, _ = w.Write([]byte("We keep breaking bash.org's mysql... eek."))
		return
	}
	re := regexp.MustCompile(`.*<p class="qt">(.*)</p>.*`)
	sma := re.FindAllStringSubmatch(page, -1)
	if len(sma) == 0 {
		fmt.Printf("Error: problem getting page...\n%s", page)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(http.ErrNoLocation.Error()))
		return
	}
	n := rand.Intn(len(sma))
	_, _ = fmt.Fprint(w, sma[n][1])
}

func insultGrab(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://randominsults.net")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	page := html.UnescapeString(string(b))
	if strings.Contains(page, "MySQL daemon") {
		w.WriteHeader(http.StatusGatewayTimeout)
		_, _ = w.Write([]byte("oof"))
		return
	}
	re := regexp.MustCompile(`<strong><i>(.*)</i></strong>`)
	sma := re.FindAllStringSubmatch(page, -1)
	if len(sma) == 0 {
		fmt.Printf("Error: problem getting page...\n%s", page)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(http.ErrNoLocation.Error()))
		return
	}
	n := rand.Intn(len(sma))
	_, _ = fmt.Fprint(w, sma[n][1])

}

func redirectRoot(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/insult", http.StatusTemporaryRedirect)
}
func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", redirectRoot)
	http.HandleFunc("/quote", quoteGrab)
	http.HandleFunc("/insult", insultGrab)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
