package controllers

import (
	"encoding/json"

	"turingbot.io/views"
)

func newFollowupCityEvent(eventName string) *views.FollowupCityEvent {
	return &views.FollowupCityEvent{
		FollowupEventInput: views.FollowupEventInput{
			Name: eventName,
		},
	}
}

// BuildFollowupCityResp ...
func BuildFollowupCityResp(name, city, when string) ([]byte, error) {
	f := newFollowupCityEvent(name)
	resp, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
