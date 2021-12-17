package main

import "time"

func StringToDate(stringDate string) time.Time {
	yourDate, _ := time.Parse("2006-01-02", stringDate)
	return yourDate
}
