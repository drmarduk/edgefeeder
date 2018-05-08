package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Edger interface {
	Date() time.Time
	Html() string
	RenewCache() error
}

type Bl struct {
	sync.Mutex
	Date     time.Time
	Html     string
	APIKey   string
	Spieltag string

	Begegnung []struct {
		ID                        string `json:"id"`
		Spieltag                  int    `json:"spieltag"`
		GeplanteAnstosszeit       string `json:"geplanteAnstosszeit"` // 20018-05-05T15:30:00+02:00
		AnstosszeitTerminiert     bool   `json:"anstosszeitTerminiert"`
		StartAnstosszeit          string `json:"startAnstosszeit"`
		EndeAnstosszeit           string `json:"endeAnstosszeit"`
		HeimTeamID                string `json:"heimTeamId"`
		AuswaertsTeamID           string `json:"auswaertsTeamId"`
		HeimTeamToreHalbzeit      int    `json:"heimTeamToreHalbzeit"`
		HeimTeamTore              int    `json:"heimTeamTore"`
		AuswaertsTeamToreHalbzeit int    `json:"auswaertsTeamToreHalbzeit"`
		AuswaertsTeamTore         int    `json:"auswaertsTeamTore"`
		Vorlaufzeit               int    `json:"vorlaufzeit"`
		Status                    string `json:"status"`
		HeimTeamNameCode          string `json:"heimTeamNameCode"`
		HeimTeamNameKurz          string `json:"heimTeamNameKurz"`
		HeimTeamNameLang          string `json:"heimTeamNameLang"`
		AuswaertsTeamNameCode     string `json:"auswaertsTeamNameCode"`
		AuswaertsTeamNameKurz     string `json:"auswaertsTeamNameKurz"`
		AuswaertsTeamNameLang     string `json:"auswaertsTeamNameLang"`

		HeimTeamURL struct {
			De string `json:"de"`
			Jp string `json:"jp"`
			En string `json:"en"`
			Es string `json:"es"`
			Pl string `json:"pl"`
		} `json:"heimTeamUrl"`
		AuswaertsTeamURL struct {
			De string `json:"de"`
			Jp string `json:"jp"`
			En string `json:"en"`
			Es string `json:"es"`
			Pl string `json:"pl"`
		} `json:"auswaertsTeamUrl"`
		MinuteOfPlay    string `json:"minute_of_play"`
		MatchDetailsURL struct {
			De string `json:"de"`
			Jp string `json:"jp"`
			En string `json:"en"`
			Es string `json:"es"`
			Pl string `json:"pl"`
		} `json:"matchDetailsUurl"`
	} `json:"begegnung"`
	WettbewerbsID   string `json:"wettbewerbsId"`
	WettbewerbsName string `json:"wettbewerbsName"`
	Saison          string `json:"saison"`
}

func NewBl() (*Bl, error) {
	b := &Bl{}

	b.Html = `<!DOCTYPE html>
	<html>
		<head>
			<meta charset="utf-8">
		</head>
		<body>
			<strong>Cache Time: $time$</strong>
			<div>$content$</div>
		</body>
	</html>`

	err := b.RefreshMetadata()
	if err != nil {
		return nil, fmt.Errorf("could not retreive API Key: %v\n", err)
	}

	err = b.RenewCache()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			time.Sleep(sleepi)
			err := b.RenewCache()
			if err != nil {
				log.Printf("could not renew cache: %v\n", err)
			}
			log.Printf("wait 30 s\n")
		}
	}()
	return b, nil
}

func (b *Bl) RenewCache() error {
	r, err := http.NewRequest("GET", fmt.Sprintf("https://api.bundesliga.com/v1/livebox-service/begegnungen/DFL-COM-000001/2017-2018/?spieltag=%s", b.Spieltag), nil)
	if err != nil {
		return err
	}

	r.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linu…) Gecko/20100101 Firefox/59.0")
	r.Header.Add("x-dflds-api-key", b.APIKey)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return fmt.Errorf("could not perform request on bl api: %v\n", err)
	}

	enc := json.NewDecoder(resp.Body)
	err = enc.Decode(b)
	if err != nil {
		return fmt.Errorf("could not decode json to bl object: %v\n", err)
	}

	b.Date = time.Now()

	return nil
}

// RefreshMetadata refreshes the apiKey and the current matchday
func (b *Bl) RefreshMetadata() error {
	resp, err := http.Get("https://www.bundesliga.com/de/bundesliga/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	src := string(_src)

	b.APIKey = getAPIKey(src)
	if b.APIKey == "" {
		return fmt.Errorf("no api key found in source")
	}

	b.Spieltag = getMatchDay(src)
	if b.Spieltag == "" {
		return fmt.Errorf("current spieltag not found")
	}
	return nil
}

func getAPIKey(src string) string {
	//apiKey: 'dwe-1-01-Gq565g-asdf-ng-55-hccnn-2cA3wUralIW',
	r := regexp.MustCompile(`apiKey: '(.*)',`)

	s := r.FindString(src)
	s = s[9 : len(s)-2]

	return s
}
func getMatchDay(src string) string {
	r := regexp.MustCompile(`'DFL-COM-000001': \d+, 'DFL-COM-000002':`)

	s := r.FindString(src)
	s = s[18:]
	s = s[:strings.Index(s, ",")]
	return s
}
