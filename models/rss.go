package models

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title       string           `xml:"title"`
	Link        string           `xml:"link"`
	Description string           `xml:"description"`
	Language    string           `xml:"language"`
	Item        []RSSChannelItem `xml:"item"`
}

type RSSChannelItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func UrlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	resp, err := httpClient.Get(url)

	if err != nil {
		return RSSFeed{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		return RSSFeed{}, err
	}
	rssFeed := RSSFeed{}

	err = xml.Unmarshal(dat, &rssFeed)

	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
