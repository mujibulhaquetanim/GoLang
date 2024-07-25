package main

import "fmt"

//passing multiple return types
func funcky() (string, string, string, bool) {
	return "js", "go", "python", true
}

func main() {
	//passing multiple return types
	lang1, lang2, lang3, bool := funcky()
	fmt.Printf("languages %t for : %s, %s, %s", bool, lang2, lang3, lang1)

}
