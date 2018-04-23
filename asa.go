package main

type AsA struct {
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
