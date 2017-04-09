package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// PowerballHandler takes requests for Powerball queries and redirects them to "https://data.ny.gov/resource/8vkr-v8vh.json"
func PowerballHandler(w http.ResponseWriter, req *http.Request) {
	var query *DateQuery
	var data []WinnerPB

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("%v", err)
	}
	if err2 := json.Unmarshal(body, &query); err2 != nil {
		log.Printf("%v", err2)
	}
	query.Reformat()
	data = getWinnersPb(query)
	body, err3 := json.Marshal(data)
	if err3 != nil {
		log.Printf("%v", err3)
	}
	w.Write(body)
}

func getWinnersPb(query *DateQuery) []WinnerPB {
	var data []WinnerPB
	var rawData []WinnerPbRaw

	if len(query.Dates) == 1 {
		api := "https://data.ny.gov/resource/8vkr-v8vh.json?draw_date=" + query.Dates[0]
		rawData = hitSocrataPbAPI(api)
		data = formatRawPbData(rawData)
	} else {
		api := "https://data.ny.gov/resource/8vkr-v8vh.json?$where="
		queryString := url.QueryEscape("draw_date between '" + query.Dates[0] + "' and '" + query.Dates[1] + "'")

		rawData = hitSocrataPbAPI(api + queryString)
		data = formatRawPbData(rawData)
	}
	return data
}

func hitSocrataPbAPI(api string) []WinnerPbRaw {
	var rawData []WinnerPbRaw

	resp, err := http.Get(api)
	if err != nil {
		log.Printf("%v", err)
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Printf("%v", err2)
	}

	if err3 := json.Unmarshal(body, &rawData); err3 != nil {
		log.Printf("%v", err3)
	}
	return rawData
}

func formatRawPbData(rawData []WinnerPbRaw) []WinnerPB {
	var data []WinnerPB

	for i, v := range rawData {
		temp := strings.Split(v.Numbers, " ")
		data = append(data, WinnerPB{})
		data[i].Powerball = temp[5]
		data[i].Numbers = strings.Join(temp[0:5], " ")
		data[i].Date = v.Date
		data[i].PowerPlay = v.PowerPlay
	}
	return data
}
