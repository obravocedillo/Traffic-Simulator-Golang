package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var gameStarted int

//Struct of the semaphores
type semaphore struct {
	//Stop = 1
	//Go = 0
	color int
	//Timer to change color
	counter int
	//Positio of the semaphore
	position int
}

//Struct of the cars
type car struct {
	speed       int
	start       int
	destination int
	//0 car is on route
	//1 car arrived
	status int
}

//Struct of the game controller
type gameController struct {
	cars               []car
	semaphores         []semaphore
	numberOfCars       int
	numberOfSemaphores int
}

//Function of semaphore
func semaphoreBehavior(currentSemaphore semaphore) {
	for {
		if currentSemaphore.counter == 20 {
			if currentSemaphore.color == 0 {
				currentSemaphore.color = 1
			} else if currentSemaphore.color == 1 {
				currentSemaphore.color = 0
			}
		} else {
			currentSemaphore.counter++
		}
	}
}

//Function of car
func carBehavior(singleCar car) {
	for {

	}
}

func main() {
	//Status of the game
	//0 Game not started
	//1 Game started
	//2 Game finished
	gameStarted = 0
	//Initialization of game controller
	mainGameController := gameController{
		numberOfCars:       5,
		numberOfSemaphores: 5,
	}
	//Options opened flag
	optionsOpen := false
	for gameStarted != 3 {
		fmt.Println("Welcome to trafic simulation, select the option you desire")
		fmt.Println("1: Start Game")
		fmt.Println("2: Change Options")
		fmt.Println("3: Exit")
		fmt.Print("Number of the option: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		option := input.Text()
		if option == "1" {
			for i := 0; i < mainGameController.numberOfCars; i++ {
				//Default values for cars
				tempCar := car{
					//Random speed from 1 to 10
					speed:  (rand.Intn(10) + 1),
					status: 0,
				}
				mainGameController.cars = append(mainGameController.cars, tempCar)
			}
			for i := 0; i < mainGameController.numberOfSemaphores; i++ {
				//Default values for semaphores
				tempSemaphore := semaphore{
					//Random speed from 1 to 10
					color:   (rand.Intn(2) + 1),
					counter: 0,
				}
				mainGameController.semaphores = append(mainGameController.semaphores, tempSemaphore)
			}
			gameStarted = 3
			fmt.Println(" ")
		} else if option == "2" {
			optionsOpen = true
			fmt.Println(" ")
			for optionsOpen == true {
				fmt.Println("Options menu, select the option you desire to change")
				fmt.Println("1: Select number of cars")
				fmt.Println("2: Select number of semaphores")
				fmt.Println("3: Return to main menu")
				fmt.Print("Number of the option: ")
				input := bufio.NewScanner(os.Stdin)
				input.Scan()
				specificOption := input.Text()
				fmt.Println(" ")
				if specificOption == "1" {
					fmt.Print("Select the number of cars: ")
					input := bufio.NewScanner(os.Stdin)
					input.Scan()
					specificOptionCars := input.Text()
					cars, _ := strconv.Atoi(specificOptionCars)
					mainGameController.numberOfCars = cars
					fmt.Println("Number of cars changed")
					fmt.Println(" ")
				} else if specificOption == "2" {
					fmt.Print("Select the number of semaphores: ")
					input := bufio.NewScanner(os.Stdin)
					input.Scan()
					specificOptionSemaphores := input.Text()
					semaphores, _ := strconv.Atoi(specificOptionSemaphores)
					mainGameController.numberOfSemaphores = semaphores
					fmt.Println("Number of semaphores changed")
					fmt.Println(" ")
				} else if specificOption == "3" {
					optionsOpen = false
					fmt.Println(" ")
				} else {
					fmt.Println("Please select a valid option")
					fmt.Println(" ")
				}
			}
		} else if option == "3" {
			fmt.Println("Thanks for playing Traffic Simulator")
			os.Exit(1)
		} else if option == "4" {

		} else {
			fmt.Print("Please select a valid option")
			fmt.Println(" ")
		}
	}

}
