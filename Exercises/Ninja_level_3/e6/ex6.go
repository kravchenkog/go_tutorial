package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	date := currentTime.Format("2006-January-02")
	if date == currentTime.Format("2006-January-02") {
		fmt.Println("Current time is", date)
	}

}
