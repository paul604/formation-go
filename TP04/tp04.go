package main

import (
	"fmt"
	"time"
)

var weekdays = []string{"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

func incr(i int) {
	i = i + 1
}

func main() {
	i := int(time.Now().Weekday())
	fmt.Printf("Today is %s\n", weekdays[i])
	incr(i)
	fmt.Printf("Tomorrow is %s\n", weekdays[i])
}
