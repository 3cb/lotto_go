package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

// MegaHandler takes requests for MegaMiilions queries and redirects them to "https://data.ny.gov/resource/h6w8-42p9.json"
func MegaHandler(w http.ResponseWriter, req *http.Request) {
	var query DateQuery
	var data []WinnerMM

	if err := body2Struct(req.Body, &query); err != nil {
		log.Println(err)
	}
	query.Reformat()

	if len(query.Dates) == 1 {
		data = getSingle(query)
	} else {
		data = getRange(query)
	}

	body, err2 := json.Marshal(data)
	if err2 != nil {
		log.Println(err2)
	}
	w.Write(body)
}

func getSingle(query DateQuery) []WinnerMM {
	var data []WinnerMM
	api := "https://data.ny.gov/resource/h6w8-42p9.json?draw_date=" + query.Dates[0]
	log.Println(api)
	resp, err := http.Get(api)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Hit Socrata API. Response Status: %#v", resp.Status)
	defer resp.Body.Close()
	if err2 := body2Struct(resp.Body, &data); err2 != nil {
		log.Printf("%+v", err2)
	}
	return data
}

func getRange(query DateQuery) []WinnerMM {
	var data []WinnerMM
	api := "https://data.ny.gov/resource/h6w8-42p9.json?$where="
	queryString := url.QueryEscape("draw_date between '" + query.Dates[0] + "' and '" + query.Dates[1] + "'")
	log.Println(api + queryString)
	resp, err := http.Get(api + queryString)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	log.Printf("Hit Socrata API. Response Status: %#v", resp.Status)
	defer resp.Body.Close()
	if err2 := body2Struct(resp.Body, &data); err2 != nil {
		log.Printf("%+v", err2)
	}
	return data
}
