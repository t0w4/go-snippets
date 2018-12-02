package main

import "fmt"

func typeJudge(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("this is integer: %v\n", v)
	case string:
		fmt.Printf("this is string: %v\n", v)
	default:
		fmt.Printf("don't know %T: %v\n", v, v)
	}
}

func printString(i interface{}) {
	fmt.Printf("this is string: %v\n", i.(string))
}

func main() {
	typeJudge(124)
	typeJudge("hogehgoe")
	typeJudge(true)

	printString("OKOKOK")
}
