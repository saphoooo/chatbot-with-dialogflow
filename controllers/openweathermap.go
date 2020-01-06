package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/saphoooo/chatbot-with-dialogflow/views"
)

//const token = "151ffc2f8114c1d959a8295d23df1596"

// Result ...
type Result struct {
	Desc    string
	TempMin float32
	TempMax float32
}

// QueryOpenweathermap ...
func QueryOpenweathermap(firstname, when, city string) ([]byte, error) {
	var o views.OpenWeather
	token := os.Getenv("OPENWEATHERMAP_TOKEN")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather/?q=" + city + "&lang=en&APPID=" + token)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return nil, errors.New("Error unmarshaling query")
	}

	return NewOpenweathermapSlackReply(firstname, when, city, o.Weather[0].Description, o.Main.TempMin, o.Main.TempMax)

}

// QueryForecast ...
func QueryForecast(firstname, when, city string) ([]byte, error) {
	var o views.OpenWeatherForecast
	token := os.Getenv("OPENWEATHERMAP_TOKEN")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast/?q=" + city + "&lang=en&APPID=" + token)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return nil, err
	}

	return NewOpenweathermapSlackReply(firstname, when, city, o.List[1].Weather[0].Description, o.List[1].Main.TempMin, o.List[1].Main.TempMax)
}
