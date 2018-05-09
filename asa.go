package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type AsA struct {
	sync.Mutex
	Date    time.Time
	Html    string
	Channel struct {
		Description string `xml:"description"`
		Title       string `xml:"title"`
		Generator   string `xml:"generator"`
		Link        string `xml:"link"`
		Item        []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Link        string `xml:"link"`
			Guid        string `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Category    string `xml:"category"`
		} `xml:"item"`
	} `xml:"channel"`
}

func NewAsA() (*AsA, error) {
	a := &AsA{}

	a.Html = `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
		</head>
		<body>
			<strong>Cache Time: $time$</strong>
			<div>$content$</div>
		</body>
	</html>`

	err := a.RenewCache()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			time.Sleep(sleepi)
			err := a.RenewCache()
			if err != nil {
				log.Printf("could not renew cache: %v\n", err)
			}
		}
	}()
	return a, nil
}

func (a *AsA) RenewCache() error {
	fmt.Printf("%s: Renew AsA Cache: ", time.Now().Format("02.01.2006 - 15:04:05"))

	t1 := time.Now()
	resp, err := http.Get("http://altschauerberganzeiger.com/rss")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// lock for writing
	a.Lock()
	defer a.Unlock()

	var cc int
	c := Counter{R: resp.Body, N: &cc}

	enc := xml.NewDecoder(c)
	err = enc.Decode(a)
	if err != nil {
		return err
	}
	a.Date = time.Now()

	took := time.Now().Sub(t1)
	fmt.Printf("read %d bytes in %.4f s\n", cc, took.Seconds())
	return nil
}
