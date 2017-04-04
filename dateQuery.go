package main

import (
	"strings"
)

type DateQuery struct {
	Dates []string `json:"dates"`
}

func (query *DateQuery) ReformatMM() {
	for i, v := range query.Dates {
		temp := strings.Split(v, "T")
		temp = append(temp[0:1], "T00:00:00.000")
		query.Dates[i] = strings.Join(temp, "")
	}
}
