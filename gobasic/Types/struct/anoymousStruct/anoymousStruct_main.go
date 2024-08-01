package main

import (
	"fmt"
)

func main() {
	// anonymous struct created and initialized in one line `struct{}{}`
	job := struct {
		title   string
		company string
		salary  int
	}{
		title:   "Software Engineer",
		company: "Google",
		salary:  100000,
	}
	fmt.Println(job)

	job.salary = 200000
	fmt.Println(job)
}
