package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/saphoooo/chatbot-with-dialogflow/controllers"
	"github.com/saphoooo/chatbot-with-dialogflow/views"
)

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("secret") != os.Getenv("TURING_BOT") {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Page not found"))
		return
	}

	var d views.QueryResult
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	firstname := strings.Title(strings.ToLower(d.QueryResult.Parameters.Firstname))
	city := d.QueryResult.Parameters.City
	when := strings.ToLower(d.QueryResult.Parameters.When)
	if when == "" {
		when = "today"
	}
	subject := d.QueryResult.Parameters.Subject
	switch action := d.QueryResult.Action; action {
	case "queryOpenweathermap":
		if firstname == "" {
			resp, err := controllers.BuildFollowupCityResp("askFirstname", city, when)
			if err != nil {
				log.Println(errors.WithMessage(err, "error in follow up sequence in queryOpenweathermap"))
			}
			controllers.JSONReply(w, resp)
		}
		switch when {
		case "today":
			resp, err := controllers.QueryOpenweathermap(firstname, when, city)
			if err != nil {
				log.Println(errors.WithMessage(err, "error in queryOpenweathermap while parsing today"))
			}
			controllers.JSONReply(w, resp)
		case "tomorrow":
			resp, err := controllers.QueryOpenweathermap(firstname, when, city)
			if err != nil {
				log.Println(errors.WithMessage(err, "error in queryOpenweathermap while parsing tomorrow"))
			}
			controllers.JSONReply(w, resp)
		case "yesterday":
			resp, err := controllers.NewSlackReply("I do not keep track of the weather for the past days.")
			if err != nil {
				log.Println(errors.WithMessage(err, "error in NewSlackReply while parsing yesterday"))
			}
			controllers.JSONReply(w, resp)
		default:
			resp, err := controllers.NewSlackReply("Currently, I can only give the weather for today or tomorrow.")
			if err != nil {
				log.Println(errors.WithMessage(err, "error in NewSlackReply with default handler"))
			}
			controllers.JSONReply(w, resp)
		}
	case "queryWikipedia":
		extract, err := controllers.QueryWikipedia(subject)
		if err != nil {
			log.Println(errors.WithMessage(err, "error in QueryWikipedia"))
		}
		err = controllers.NewWikipediaSlackReply(w, subject, *extract)
	case "queryUSPresident":
		err := controllers.NewUSPresidentQuery(w)
		if err != nil {
			log.Println(errors.WithMessage(err, "error in queryUSPresident"))
		}
	default:
		log.Printf("Action %s is not configured\n", action)
	}

}

func main() {
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")
	r := mux.NewRouter()
	r.HandleFunc("/hL3DO1QXdlw2liRkwQZ84JagiT392j6xBChZ67U7", webhook).Methods("POST")
	err := http.ListenAndServeTLS(":8080", certFile, keyFile, r)
	if err != nil {
		panic(err)
	}
}
