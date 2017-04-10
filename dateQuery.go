package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// DateQuery contains date(s) for Socrata API query
type DateQuery struct {
	Dates []string `json:"dates"`
}

// Reformat reformats date object so it suits Socrata API
func (query *DateQuery) Reformat() {
	for i, v := range query.Dates {
		temp := strings.Split(v, "T")
		temp = append(temp[0:1], "T00:00:00.000")
		query.Dates[i] = strings.Join(temp, "")
	}
}

func body2Struct(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err2 := json.Unmarshal(body, v); err2 != nil {
		return err2
	}
	return nil
}
