package main

import "fmt"

func main() {
	testingcallback(1, func(i int) int {
		return i * 20
	})
}

func testingcallback(num int, callback func(i int) int) {
	testing := callback(num)

	fmt.Println(testing)
}
