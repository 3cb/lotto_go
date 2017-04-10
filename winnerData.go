package main

// WinnerMM holds data for winning MegaMillions draws
type WinnerMM struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	MegaBall  string `json:"mega_ball"`
	Megaplier string `json:"multiplier"`
}

// WinnerPbRaw holds unformatted data for winning Powerball draws
type WinnerPbRaw struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	PowerPlay string `json:"multiplier"`
}

// WinnerPB holds formatted data for winning Powerball draws
type WinnerPB struct {
	Date      string `json:"draw_date"`
	Numbers   string `json:"winning_numbers"`
	Powerball string `json:"powerball"`
	PowerPlay string `json:"multiplier"`
}
