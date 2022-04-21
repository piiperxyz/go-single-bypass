package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	hour := t.Hour()
	minute := t.Minute()
	day := t.Day()
	//hour1 := strconv.Itoa(hour)
	pass := strconv.Itoa(hour) + strconv.Itoa(minute) + strconv.Itoa(day)
	fmt.Printf("%s", pass)
}
