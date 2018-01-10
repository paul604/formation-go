package main

import (
    "fmt"
    "github.com/paul604/formation-go/TP02/golaunch"
)

func main() {
    d := golaunch.GetDaysSinceGoLaunch()
    fmt.Printf("Hello it's %d days since Go Launch\n", int(d))
}

//func getDaysSinceGoLaunch() float64 {
//    t0 := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
//    t1 := time.Now()
//    duration := t1.Sub(t0)
//    return duration.Hours() / 24
//}
