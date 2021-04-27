package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func quoteGrab(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	resp, _ := http.Get("http://bash.org/?random1")
	b, _ := ioutil.ReadAll(resp.Body)
	page := html.UnescapeString(string(b))
	re := regexp.MustCompile(`.*<p class="qt">(.*)</p>.*`)
	sma := re.FindAllStringSubmatch(page, -1)
	n := rand.Intn(len(sma))
	_, _ = fmt.Fprint(w, sma[n][1])
}
func main()  {
	http.HandleFunc("/", quoteGrab)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return 
	}
}
