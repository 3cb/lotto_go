package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// MegaHandler takes requests for MegaMiilions queries and redirects them to "https://data.ny.gov/resource/h6w8-42p9.json"
func MegaHandler(w http.ResponseWriter, req *http.Request) {
	var query *DateQuery
	var data []WinnerMM

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	if err2 := json.Unmarshal(body, &query); err2 != nil {
		log.Println(err2)
	}
	query.Reformat()

	if len(query.Dates) == 1 {
		data = getSingle(query)
	} else {
		data = getRange(query)
	}

	body, err3 := json.Marshal(data)
	if err3 != nil {
		log.Println(err3)
	}
	w.Write(body)
}

func getSingle(query *DateQuery) []WinnerMM {
	var data []WinnerMM
	api := "https://data.ny.gov/resource/h6w8-42p9.json?draw_date=" + query.Dates[0]
	log.Println(api)
	resp, err := http.Get(api)
	log.Printf("Hit Socrata API. Response Status: %#v", resp.Status)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println(err2)
	}
	if err3 := json.Unmarshal(body, &data); err3 != nil {
		log.Println(err3)
	}
	return data
}

func getRange(query *DateQuery) []WinnerMM {
	var data []WinnerMM
	api := "https://data.ny.gov/resource/h6w8-42p9.json?$where="
	queryString := url.QueryEscape("draw_date between '" + query.Dates[0] + "' and '" + query.Dates[1] + "'")
	resp, err := http.Get(api + queryString)
	log.Printf("Hit Socrata API. Response Status: %#v", resp.Status)
	if err != nil {
		log.Printf("err: %+v", err)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Printf("err2: %+v", err2)
	}
	if err3 := json.Unmarshal(body, &data); err3 != nil {
		log.Printf("err3: %#v", err3)
	}
	return data
}
