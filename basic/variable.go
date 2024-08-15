package main

import "fmt"

func main() {
	//	testingcallback(1, func(i int) int {
	//		return i * 20
	//	})
	data := []string{"Andi", "Ana", "Ali", "Budi", "Caca", "Darius"}
	function := returnFunction(data)

	isExist := function("Ands")

	fmt.Print(isExist)

}

func testingcallback(num int, callback func(i int) int) {
	testing := callback(num)

	fmt.Println(testing)
}

func returnFunction(data []string) func(string) bool {
	return func(key string) bool {
		isExist := false
		for i := 0; i < len(data); i++ {
			if data[i] == key {
				isExist = true
				break
			}
		}
		return isExist
	}
}
