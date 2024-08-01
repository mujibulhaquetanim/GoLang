package main

import "fmt"

type Vehicle interface {
	start()
	stop()
}

type Car struct {
	wheelNum int
}

func (c Car) start() {
	fmt.Println("The car has Started")
}

func (c Car) stop() {
	fmt.Println("The car has stopped")
}

type MotorBike struct {
	wheelNum int
}

func (m MotorBike) start() {
	fmt.Printf("The MotorBike has Started with %d", m.wheelNum)
}

func (m MotorBike) stop() {
	fmt.Printf("\nThe MotorBike has stopped")
}

func operateVehicle(v Vehicle) {
	v.start()
	v.stop()
}

func main() {
	// myCar:= Car{}
	myBike := MotorBike{wheelNum: 2}
	//without inserting hardcode code, we can inject any instance that meets vehicle contract.
	//my car instance has been injected to the normal function that takes an Vehicle instance
	// operateVehicle(myCar)
	operateVehicle(myBike)
}
