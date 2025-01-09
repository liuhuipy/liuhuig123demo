package model

import "gorm.io/gorm"

type HotelInfo struct {
	gorm.Model
	Name             string
	Star             int32
	Price            float64
	PriceBeforeTaxes string
	CheckInDate      string
	CheckOutDate     string
	Guests           string
}
