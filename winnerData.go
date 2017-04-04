package main

type Winner struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	MegaBall  string `json:"mega_ball"`
	Megaplier string `json:"multiplier"`
}
