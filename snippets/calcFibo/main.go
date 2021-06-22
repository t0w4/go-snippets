package main

import "fmt"

func main() {
	fmt.Println(fibo2(45))
}

// ex) 0 1 1 2 3
func fibo1(num int) int {
	if num < 2 {
		return num
	}
	return fibo1(num-1) + fibo1(num-2)
}

var memo = make(map[int]int)

// ex) 0 1 1 2 3
func fibo2(num int) int {
	if res, ok := memo[num]; ok {
		return res
	}
	if num < 2 {
		return num
	}
	r := fibo2(num-1) + fibo2(num-2)
	memo[num] = r
	return r
}
