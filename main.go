package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/asa", asa)

	log.Fatalln(http.ListenAndServe(":9001", nil))
}

var (
	reqError string = `{"error": true, "msg": "%s"}`
)

func asa(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://altschauerberganzeiger.com/rss")
	if err != nil {
		fmt.Fprintf(w, reqError, err)
		return
	}

	defer resp.Body.Close()
	// TODO: do html stuff
	enc := xml.NewDecoder(resp.Body)
	var data AsA = AsA{}
	err = enc.Decode(&data)
	if err != nil {
		log.Fatalf("XML error: %v\n", err)
	}

	x := unHtml(data.Channel.Item[0].Description)

	fmt.Fprintf(w, x)
	fmt.Fprintf(w, "<br><br>%s", data.Channel.Item[0].Description)
}

func unHtml(src string) string {
	return html.UnescapeString(src)
}
