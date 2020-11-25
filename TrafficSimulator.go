package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var gameStarted int 
var mainGameController gameController

//Struct of the semaphores
type semaphore struct {
	//Stop = 1
	//Go = 0
	color int
	//Timer to change color
	counter int
	//Positio of the semaphore
	positionX float64
	positionY float64
	image *ebiten.Image
	options ebiten.DrawImageOptions
}

//Struct of the cars
type car struct {
	speed       float64
	start       int
	destination int
	//0 car is on route
	//1 car arrived
	status int
	positionX float64
	positionY float64
	image *ebiten.Image
	options ebiten.DrawImageOptions
}

//Struct of the game controller
type gameController struct {
	cars               []car
	semaphores         []semaphore
	numberOfCars       int
	numberOfSemaphores int
	screenWidth int 
	screenHeight int
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
func carBehavior(carIndex int) {
	for {
		//Movement of the cars
		mainGameController.cars[carIndex].positionX = mainGameController.cars[carIndex].speed;
	}
}

func drawBoard(screen *ebiten.Image){
	//Every sprite is 80*80 for moving the road just add or remove 80
	roadMain, _, _ := ebitenutil.NewImageFromFile("roads/roadNEWS.png", ebiten.FilterDefault)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)-40)
	screen.DrawImage(roadMain,op)

	road1, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(mainGameController.screenWidth/2)+40, float64(mainGameController.screenHeight/2)-40)
	screen.DrawImage(road1,op2)

	road2, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Translate(float64(mainGameController.screenWidth/2)-120, float64(mainGameController.screenHeight/2)-40)
	screen.DrawImage(road2,op3)

	road3, _, _ := ebitenutil.NewImageFromFile("roads/roadNS.png", ebiten.FilterDefault)
	op4 := &ebiten.DrawImageOptions{}
	op4.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)+40)
	screen.DrawImage(road3,op4)

	road4, _, _ := ebitenutil.NewImageFromFile("roads/roadNS.png", ebiten.FilterDefault)
	op5 := &ebiten.DrawImageOptions{}
	op5.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)-120)
	screen.DrawImage(road4,op5)

	road5, _, _ := ebitenutil.NewImageFromFile("roads/roadSE.png", ebiten.FilterDefault)
	op6 := &ebiten.DrawImageOptions{}
	op6.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)-200)
	screen.DrawImage(road5,op6)

	road6, _, _ := ebitenutil.NewImageFromFile("roads/roadSE.png", ebiten.FilterDefault)
	op7 := &ebiten.DrawImageOptions{}
	op7.GeoM.Translate(float64(mainGameController.screenWidth/2)-200, float64(mainGameController.screenHeight/2)-40)
	screen.DrawImage(road6,op7)

	road8, _, _ := ebitenutil.NewImageFromFile("roads/roadSW.png", ebiten.FilterDefault)
	op9 := &ebiten.DrawImageOptions{}
	op9.GeoM.Translate(float64(mainGameController.screenWidth/2)+120, float64(mainGameController.screenHeight/2)-40)
	screen.DrawImage(road8,op9)

	road10, _, _ := ebitenutil.NewImageFromFile("roads/roadNS.png", ebiten.FilterDefault)
	op11 := &ebiten.DrawImageOptions{}
	op11.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)+120)
	screen.DrawImage(road10,op11)

	road11, _, _ := ebitenutil.NewImageFromFile("roads/roadNW.png", ebiten.FilterDefault)
	op12 := &ebiten.DrawImageOptions{}
	op12.GeoM.Translate(float64(mainGameController.screenWidth/2)+200, float64(mainGameController.screenHeight/2)-200)
	screen.DrawImage(road11,op12)

	road13, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op14 := &ebiten.DrawImageOptions{}
	op14.GeoM.Translate(float64(mainGameController.screenWidth/2)+120, float64(mainGameController.screenHeight/2)-200)
	screen.DrawImage(road13,op14)

	road14, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op15 := &ebiten.DrawImageOptions{}
	op15.GeoM.Translate(float64(mainGameController.screenWidth/2)+40, float64(mainGameController.screenHeight/2)-200)
	screen.DrawImage(road14,op15)

	road15, _, _ := ebitenutil.NewImageFromFile("roads/roadNS.png", ebiten.FilterDefault)
	op16 := &ebiten.DrawImageOptions{}
	op16.GeoM.Translate(float64(mainGameController.screenWidth/2)+120, float64(mainGameController.screenHeight/2)+40)
	screen.DrawImage(road15,op16)

	road16, _, _ := ebitenutil.NewImageFromFile("roads/roadNE.png", ebiten.FilterDefault)
	op17 := &ebiten.DrawImageOptions{}
	op17.GeoM.Translate(float64(mainGameController.screenWidth/2)+120, float64(mainGameController.screenHeight/2)+120)
	screen.DrawImage(road16,op17)

	road17, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op18 := &ebiten.DrawImageOptions{}
	op18.GeoM.Translate(float64(mainGameController.screenWidth/2)+200, float64(mainGameController.screenHeight/2)+200)
	screen.DrawImage(road17,op18)

	road19, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op20 := &ebiten.DrawImageOptions{}
	op20.GeoM.Translate(float64(mainGameController.screenWidth/2)+200, float64(mainGameController.screenHeight/2)+120)
	screen.DrawImage(road19,op20)

	road20, _, _ := ebitenutil.NewImageFromFile("roads/roadNW.png", ebiten.FilterDefault)
	op21 := &ebiten.DrawImageOptions{}
	op21.GeoM.Translate(float64(mainGameController.screenWidth/2)-200, float64(mainGameController.screenHeight/2)+40)
	screen.DrawImage(road20,op21)

	road21, _, _ := ebitenutil.NewImageFromFile("roads/roadEW.png", ebiten.FilterDefault)
	op22 := &ebiten.DrawImageOptions{}
	op22.GeoM.Translate(float64(mainGameController.screenWidth/2)-280, float64(mainGameController.screenHeight/2)+40)
	screen.DrawImage(road21,op22)


}

//Game Loo
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	//This part of the code draws the static road
	drawBoard(screen)
	//This part of the code draws the static road

	//This part of the code draws all the cars in the screen and update the position of the cars
	for i := 0; i < mainGameController.numberOfCars; i++ {
		mainGameController.cars[i].options.GeoM.Translate(mainGameController.cars[i].positionX, mainGameController.cars[i].positionY)
		screen.DrawImage(mainGameController.cars[i].image, &mainGameController.cars[i].options)
	}
	return nil
}

func main() {
	//Status of the game
	//0 Game not started
	//1 Game started
	//2 Game finished
	gameStarted = 0
	//Initialization of game controller
	mainGameController = gameController{
		numberOfCars:       5,
		numberOfSemaphores: 5,
		screenWidth: 620,	
		screenHeight: 400,

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
				currentImage, _, err := ebitenutil.NewImageFromFile("roads/car.png", ebiten.FilterDefault)
				if err != nil {
					log.Fatal(err)
				}
				tempCar := car{
					//Random speed from 1 to 10
					speed: (rand.Float64() * (10-1)),
					status: 0,
					//Image of the car
					image: currentImage,
					positionX: 0.0,
					positionY: 0.0,
					//Draw Options of the car
					options: ebiten.DrawImageOptions{},
				}
				mainGameController.cars = append(mainGameController.cars, tempCar)
				go carBehavior(i)
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
			if err := ebiten.Run(update, mainGameController.screenWidth, mainGameController.screenHeight, 2, "Traffic Simulator"); err != nil {
				log.Fatal(err)
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
