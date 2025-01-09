package dto

import "net/http"

type Task struct {
	TaskId  string      `json:"task_id"`
	Name    string      `json:"name"`
	Url     string      `json:"url"`
	Headers http.Header `json:"headers"`
}

type TasksJson struct {
	Tasks []*Task `json:"tasks"`
}

type Hotel struct {
	HotelName        string  `json:"hotel_name"`
	Star             int32   `json:"star"`
	Price            float64 `json:"price"`
	PriceBeforeTaxes string  `json:"price_before_taxes"`
	CheckInDate      string  `json:"check_in_date"`
	CheckOutDate     string  `json:"check_out_date"`
	Guests           string  `json:"guests"`
}
