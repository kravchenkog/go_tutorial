package main

import (
	"fmt"
	"time"
)

func main() {
	meetingTime := time.Now().AddDate(0, 0, 1).Format("2006-January-02")
	resultPrinter(meetingTime)

}

func resultPrinter(meetingTime string) {
	today := time.Now()
	if today.Format("2006-January-02") == meetingTime {
		fmt.Println("The meeting will be Today: ", meetingTime)

	} else if today.AddDate(0, 0, 1).Format("2006-January-02") == meetingTime {
		fmt.Println("The meeting will be tomorrow: ", meetingTime)

	} else {
		fmt.Println("The date of the meeting is: ", meetingTime)
	}

}
