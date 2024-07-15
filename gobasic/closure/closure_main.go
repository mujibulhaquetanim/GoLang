package main

import "fmt"

func activateGiftCard() func(int) int {
	amount := 100

	debitFunc := func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
	return debitFunc
}

func main() {

	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()
	useGiftCard3 := activateGiftCard()

	//each call will have a new closure, the return func will hold the value of amount even after the activateGiftCard call
	fmt.Println(useGiftCard1(5)) // 95
	fmt.Println(useGiftCard2(10)) // 90
	fmt.Println(useGiftCard3(15)) // 85

}
