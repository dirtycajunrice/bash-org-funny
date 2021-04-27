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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := html.UnescapeString(string(b))
	if strings.Contains(page, "Sorry, the MySQL daemon appears to be down.1Error: problem getting page...") {
		http.Error(w, "We keep breaking bash.org's mysql... eek.", http.StatusGatewayTimeout)
	}
	re := regexp.MustCompile(`.*<p class="qt">(.*)</p>.*`)
	sma := re.FindAllStringSubmatch(page, -1)
	if len(sma) == 0 {
		fmt.Printf("Error: problem getting page...\n%s", page)
		http.Error(w, http.ErrNoLocation.Error(), http.StatusInternalServerError)
		return
	}
	n := rand.Intn(len(sma))
	_, _ = fmt.Fprint(w, sma[n][1])
}
func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", quoteGrab)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
