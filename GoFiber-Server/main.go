package main

import("fmt","github.com/gofiber/fiber/v2")

func main(){
	fmt.Println("welcome to GoFiber Server")
	app := fiber.New()
	app.listen(3000)
}