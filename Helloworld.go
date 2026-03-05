package main

import "fmt"

func zoo(animal string, sound string) {
	fmt.Printf("The %s goes %s\n", animal, sound)
}

func add(x int, y int) int {
	x += 10
	y += 20
	return x + y
}

func addptr(x *int, y *int) int {

	*x += 10
	*y += 20
	return *x + *y
}

func main() {

	/*	var nums [50]int
		for i := 0; i < 50; i++ {
			nums[i] = i
		}
		fmt.Println(nums)
	*/
	//zoo("dog", "woof")

	/*	balance := [...]float64{100.0, 200.0, 300.0}
		fmt.Println(balance)
	*/
	baLance := [][]float64{
		{100.0, 200.0, 300.0},
		{400.0, 500.0, 600.0},
	}

	for i, row := range baLance {
		fmt.Printf("Row %d: %v\n", i, row)
	}

}
