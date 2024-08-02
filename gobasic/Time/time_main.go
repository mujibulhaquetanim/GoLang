package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("time now: ", time.Now())

	fmt.Println("time now: ", time.Now().UnixNano())

	fmt.Println("time now: ", time.Now().UnixMicro())

	fmt.Println("time now: ", time.Now().Unix())

	fmt.Println("time now: ", time.Now().Format("01/02/2006 03:04:05 PM"))

	fmt.Println("time now: ", time.Now().Add(time.Minute*5))

	fmt.Println("time now: ", time.Now().Add(time.Millisecond*5))

	fmt.Println("time now: ", time.Now().Add(time.Second*5))

	fmt.Println("time now: ", time.Now().Add(time.Hour*5))

	fmt.Println("time now: ", time.Now().Add(time.Since(time.Now())))

	//time since january, 2024
	fmt.Println("time now: ", time.Now().Add(time.Since(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))))

	fmt.Println("time now: ", time.Now().Add(time.Until(time.Now())))

	fmt.Println("time now: ", time.Now().Add(time.Duration(5)*time.Minute))

	fmt.Println("time now: ", time.After(time.Millisecond*5))

	fmt.Println("time now: ", time.Now().After(time.Now()))

	fmt.Println("time now : ", time.Now().Add(time.Millisecond*5).Before(time.Now()))

	fmt.Println("time now: ", time.Now().Add(time.Millisecond*5).Format("01/02/2006 03:04:05 PM"))

}
