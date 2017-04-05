package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/api/megamillions", MegaHandler)
	r.HandleFunc("/api/powerball", PowerballHandler)

	n := negroni.New(negroni.NewStatic(http.Dir("./")))
	n.UseHandler(r)
	log.Fatal(http.ListenAndServe(":3000", n))
}

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(200)
// 	io.WriteString(w, "Welcome to the Homepage!!")
// }

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
