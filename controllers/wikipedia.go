package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/saphoooo/chatbot-with-dialogflow/views"
)

// NewWikiSummary ...
func NewWikiSummary() *views.WikiSummary {
	return &views.WikiSummary{}
}

// QueryWikipedia ...
func QueryWikipedia(name string) (*views.WikiSummary, error) {
	s := NewWikiSummary()
	nameArray := strings.Split(name, " ")
	nameUnderscore := strings.Join(nameArray, "_")
	resp, err := http.Get("https://en.wikipedia.org/api/rest_v1/page/summary/" + nameUnderscore)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
