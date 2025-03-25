package service

import (
	"fmt"
	"time"
)

type TwillioAntiSpam struct {
	AddOns struct {
		Code    any `json:"code"`
		Message any `json:"message"`
		Results struct {
			NomoroboSpamscore struct {
				Code       any    `json:"code"`
				Message    any    `json:"message"`
				RequestSid string `json:"request_sid"`
				Result     struct {
					Message       string `json:"message"`
					NeighborScore int    `json:"neighbor_score"`
					Score         int    `json:"score"`
					Status        string `json:"status"`
				} `json:"result"`
				Status string `json:"status"`
			} `json:"nomorobo_spamscore"`
		} `json:"results"`
		Status string `json:"status"`
	} `json:"add_ons"`
	CallerName     any    `json:"caller_name"`
	Carrier        any    `json:"carrier"`
	CountryCode    string `json:"country_code"`
	NationalFormat string `json:"national_format"`
	PhoneNumber    string `json:"phone_number"`
	URL            string `json:"url"`
}

func FixDateTime(t time.Time) time.Time {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading Time Zone:", err)
		return t
	}
	if t.Location() != time.UTC {
		return t
	}

	tInCST := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
	return tInCST
}
