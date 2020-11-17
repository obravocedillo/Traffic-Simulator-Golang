package main

import (
	"bufio"
	"fmt"
	"os"
)

var gameStarted int

//Struct of the semaphores
type semaphore struct {
	color    int
	counter  int
	position int
}

//Struct of the cars
type car struct {
	speed       int
	start       int
	destination int
	status      int
}

//Struct of the game controller
type gameController struct {
	cars               []car
	semaphores         []semaphore
	numberOfCars       int
	numberOfSemaphores int
}

//Function of semaphore
func semaphoreBehavior() {
	for {

	}
}

//Function of car
func carBehavior(singleCar car) {
	for {
		singleCar.speed = 10
	}
}

func main() {
	//Status of the game
	//0 Game not started
	//1 Game started
	//2 Game finished
	gameStarted = 0
	//Initialization of game controller
	mainGameController := gameController{}
	//Default values for game controller
	mainGameController.numberOfCars = 5
	mainGameController.numberOfSemaphores = 5
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
