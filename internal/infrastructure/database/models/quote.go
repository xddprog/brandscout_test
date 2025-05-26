package models


type Quote struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type CreateQuote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}


