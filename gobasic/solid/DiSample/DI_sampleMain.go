package main

import "fmt"

type Vehicle interface {
	start()
	stop()
}

type Car struct{}

func (c Car) start(){
	fmt.Println("The car has Started")
}

func (c Car) stop(){
	fmt.Println("The car has stopped")
}

type MotorBike struct{}

func (c MotorBike) start(){
	fmt.Println("The MotorBike has Started")
}

func (c MotorBike) stop(){
	fmt.Println("The MotorBike has stopped")
}

func operateVehicle(v Vehicle){
	v.start()
	v.stop()
}

func main(){
	// myCar:= Car{}
	myBike:= MotorBike{}
	//without inserting hardcode code, we can inject any instance that meets vehicle contract.
	//my car instance has been injected to the normal function that takes an Vehicle instance
	// operateVehicle(myCar)
	operateVehicle(myBike)
}