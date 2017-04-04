package main

type WinnerMM struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	MegaBall  string `json:"mega_ball"`
	Megaplier string `json:"multiplier"`
}

type WinnerPbRaw struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	PowerPlay string `json:"multiplier"`
}

type WinnerPB struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	Powerball string `json:"powerball"`
	PowerPlay string `json:"multiplier"`
}
