package main

import ("fmt"
 "context"
 "time")

func main() {
	//create a background context
	ctx := context.Background()
	fmt.Println(ctx)

	//create a context with a deadline
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	fmt.Println(ctx)

}
