package main

import "fmt"

func main() {
	var distance int
	fmt.Scan(&distance)

	var time int
	fmt.Scan(&time)

	var velocity int
	fmt.Scan(&velocity)

	var fuel int
	fmt.Scan(&fuel)

	var fuelConsumption int
	fmt.Scan(&fuelConsumption)

	if velocity*time < distance {
		fmt.Println("Failure: Not enought time!")
	} else if fuel*fuelConsumption < distance {
		fmt.Println("Failure: Not enough fuel")
	} else {
		fmt.Println("Success: Welcome to Mars")
	}
}
