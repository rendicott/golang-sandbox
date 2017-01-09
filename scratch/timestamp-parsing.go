package main

import (
"fmt"
"time"
)

func main() {
    fmt.Println(time.Now())
    // use the strange parsing format for dates
    //  more info here: https://pauladamsmith.com/blog/2011/05/go_time.html
    const dateFormat = "2006-01-02T15:04:05-07:00"
    // now use that format definition to actually parse your date
    t, err := time.Parse(dateFormat, "2012-09-06T08:11:48-05:00")
        if err != nil {
            fmt.Println("parse error: ", err.Error())
        }
        // and then print the format out in ANSIC format
        fmt.Println(t.Format(time.ANSIC))
}