package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	listen   string = "localhost:9001"
	reqError string = `{"error": true, "msg": "%s"}`
	asaCache *AsA
	blCache  *Bl

	sleepi time.Duration = 24 * time.Hour
)

func main() {
	var err error
	asaCache, err = NewAsA()
	if err != nil {
		log.Printf("could not initialize asa Cache, wait for next round: %v\n", err)
	}

	blCache, err = NewBl()
	if err != nil {
		log.Printf("could not initialize asa Cache, wait for next round: %v\n", err)
	}

	http.HandleFunc("/asa", asa)
	http.HandleFunc("/bl", bl)
	log.Printf("listen on http://%s/\n", listen)
	log.Fatalln(http.ListenAndServe(listen, nil))
}

func asa(w http.ResponseWriter, r *http.Request) {
	h := strings.Replace(asaCache.Html, "$time$", asaCache.Date.Format("02.01.2006 - 15:04:05"), 1)
	h = strings.Replace(h, "$content$", asaCache.Channel.Item[0].Description, 1)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, h)
}

func bl(w http.ResponseWriter, r *http.Request) {
	h := strings.Replace(blCache.Html, "$time$", blCache.Date.Format("02.01.2006 - 15:04:05"), 1)

	buf := bytes.NewBufferString("")

	for _, m := range blCache.Begegnung {
		buf.WriteString(fmt.Sprintf("<p>%s %d - %d %s</p>", m.HeimTeamNameLang, m.HeimTeamTore, m.AuswaertsTeamTore, m.AuswaertsTeamNameLang))
	}

	h = strings.Replace(h, "$content$", buf.String(), 1)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, h)
}
