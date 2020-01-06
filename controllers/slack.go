package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/saphoooo/chatbot-with-dialogflow/views"
)

// NewUSPresidentQuery ...
func NewUSPresidentQuery(w http.ResponseWriter) error {
	query := views.NewSlack{
		Payload: views.Payload{
			Slack: views.Slack{
				Text: "President names",
				Attachments: []views.Attachments{
					views.Attachments{
						Color:      "#c2a2d3",
						Text:       "About which president do you want informations?",
						CallbackID: "test_bot_sucks",
						Actions: []views.Actions{
							views.Actions{
								Name:  "selectPresident",
								Text:  "Trump",
								Value: "Who is Donald Trump",
								Type:  "button",
								Style: "danger",
							},
							views.Actions{
								Name:  "selectPresident",
								Text:  "Obama",
								Value: "Who is Barack Obama",
								Type:  "button",
							},
							views.Actions{
								Name:  "selectPresident",
								Text:  "Clinton",
								Type:  "button",
								Value: "Who is Bill Clinton",
								Style: "primary",
							},
						},
					},
				},
			},
		},
	}
	resp, err := json.Marshal(query)
	if err != nil {
		return errors.New("Failed to marshal response")
	}
	JSONReply(w, resp)
	return nil
}

// NewWikipediaSlackReply ...
func NewWikipediaSlackReply(w http.ResponseWriter, subject string, summary views.WikiSummary) error {
	reply := views.NewSlack{
		Payload: views.Payload{
			Slack: views.Slack{
				Attachments: []views.Attachments{
					views.Attachments{
						Color:    "#c2a2d3",
						Title:    subject,
						Text:     summary.Extract,
						ThumbURL: summary.Thumbnail.Source,
						Footer:   "Extract from Wikipedia",
					},
				},
			},
		},
	}
	resp, err := json.Marshal(reply)
	if err != nil {
		return errors.New("Failed to marshal response")
	}
	JSONReply(w, resp)
	return nil
}

// NewOpenweathermapSlackReply ...
func NewOpenweathermapSlackReply(firstname, when, city, desc string, tempMin, tempMax float32) ([]byte, error) {
	tempMinCelsius := kelvinToCelsius(tempMin)
	tempMaxCelsius := kelvinToCelsius(tempMax)
	min := fmt.Sprintf("%.0f", tempMinCelsius)
	max := fmt.Sprintf("%.0f", tempMaxCelsius)
	text := "Hi " + firstname + ", " + when + " in " + city + " the weather is: " + desc + " with temperatures between " + min + " and " + max + " Celsius degrees."
	return NewSlackReply(text)
}

// NewSlackReply ...
func NewSlackReply(message string) ([]byte, error) {
	reply := &views.NewSlack{
		FulfillmentText: message,
		Payload: views.Payload{
			Slack: views.Slack{
				Text: message,
			},
		},
	}
	resp, err := json.Marshal(reply)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
