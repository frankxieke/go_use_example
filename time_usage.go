package main
import (
    "time"
    "fmt"
)

func main() {
    str := "2018-10-01 03:01:00"
    fmt.Println("1. str: ", str)
    TIME_LAYOUT:="2006-01-02 15:04:05"
    t, _ := time.Parse(TIME_LAYOUT, str)
    fmt.Println("2. Parse time: ", t)
    t2:= t.Format("2006-01-02")
    fmt.Println("2. Format time: ", t2)
}
