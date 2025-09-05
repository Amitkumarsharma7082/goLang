package main

import (
	"fmt"
	"time"
)

func main() {
	// YYYY-MM-DD HH:MM:SS

	currentTime := time.Now()
	fmt.Println("Current Time :", currentTime) // 2025-09-05 15:43:14.9941 +0530 IST m=+0.000240139
	fmt.Printf("Type is : %T\n", currentTime)  // time.Time

	formatted := currentTime.Format("02-01-2006, Monday, 03:04:05") // dd-mm-yyyy(not allowed add specific 02-01-2006)
	fmt.Println("Format date :", formatted)

	layout := "02/01/2006"
	dateStr := "25/11/2025"
	formatted_time, _ := time.Parse(layout, dateStr)
	fmt.Println("formatted time :", formatted_time) // if it is mismatched : 0001-01-01 00:00:00 +0000 UTC

	// Add 24 hours in current date
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new_date :", new_date)
	formatted_new_date := new_date.Format("Monday")
	fmt.Println("format new date :", formatted_new_date)

}
