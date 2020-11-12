package main

import (
	"bufio"
	"fmt"
	"os"
)

var gameStarted int

type semaphore struct {
	color    int
	counter  int
	position int
}

type car struct {
	speed       int
	start       int
	destination int
	status      int
}

type gameController struct {
	cars               []car
	semaphores         []semaphore
	numberOfCars       int
	numberOfSemaphores int
}

func main() {
	gameStarted = 0
	for gameStarted != 3 {
		fmt.Println("Welcome to trafic simulation, select the option you desire")
		fmt.Println("1: Start Game")
		fmt.Println("2: Change Options")
		fmt.Println("3: Exit")
		fmt.Print("Option: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		option := input.Text()
		if option == "3" {
			gameStarted = 3
		}
	}

}
