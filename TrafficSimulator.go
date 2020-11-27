package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"image/color"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"strconv"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"time"
)

var gameStarted int 
var mainGameController gameController
var normalFont font.Face
const dpi = 60

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
	options *ebiten.DrawImageOptions
}

//Struct of the cars
type car struct {
	speedX       float64
	speedY       float64
	start       int
	destination int
	//0 car is on route
	//1 car arrived
	status int
	positionX float64
	positionY float64
	rotationType int
	image *ebiten.Image
	options ebiten.DrawImageOptions
}

type traficLightController struct {
	lightsImage        *ebiten.Image
	lightsOptions      *ebiten.DrawImageOptions
	positionX int
	positionY int
	rotationType int
}

//Struct of the game controller
type gameController struct {
	cars               []car
	semaphores         []semaphore
	traficLightControllers []traficLightController
	numberOfCars       int
	numberOfSemaphores int
	screenWidth int 
	screenHeight int
	startingPositionsX []float64
	startingPositionsY []float64
}

//Function of semaphore
func semaphoreBehavior(carIndex int) {
	for {
		//Changin of color
		if mainGameController.semaphores[carIndex].counter == 8 {
			if mainGameController.semaphores[carIndex].color == 0 {
				mainGameController.semaphores[carIndex].color = 1
			} else if mainGameController.semaphores[carIndex].color == 1 {
				mainGameController.semaphores[carIndex].color = 0
			}
			mainGameController.semaphores[carIndex].counter = 0
		} else {
			mainGameController.semaphores[carIndex].counter++
			time.Sleep(1 * time.Second)
		}
	}
}

//Function of car
func carBehavior(carIndex int) {
	for {
		time.Sleep(1 * time.Second)
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

	lights1, _, _ := ebitenutil.NewImageFromFile("roads/light.png", ebiten.FilterDefault)
	opl1 := &ebiten.DrawImageOptions{}
	opl1.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)-70)
	traficLightController1 := traficLightController{
		lightsImage: lights1,
		lightsOptions: opl1,
		positionX: (mainGameController.screenWidth/2)-40,
		positionY: (mainGameController.screenHeight/2)-70,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController1)
	mainGameController.semaphores[0].options = opl1
	mainGameController.semaphores[0].image = lights1

	lights2, _, _ := ebitenutil.NewImageFromFile("roads/light.png", ebiten.FilterDefault)
	opl2 := &ebiten.DrawImageOptions{}
	opl2.GeoM.Translate(float64(mainGameController.screenWidth/2)+125, float64(mainGameController.screenHeight/2)+40)
	traficLightController2 := traficLightController{
		lightsImage: lights2,
		lightsOptions: opl2,
		positionX: (mainGameController.screenWidth/2)+125,
		positionY: (mainGameController.screenHeight/2)+40,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController2)
	mainGameController.semaphores[1].options = opl2
	mainGameController.semaphores[1].image = lights2
	

	lights3, _, _ := ebitenutil.NewImageFromFile("roads/light.png", ebiten.FilterDefault)
	opl3 := &ebiten.DrawImageOptions{}
	opl3.GeoM.Translate(float64(mainGameController.screenWidth/2)-40, float64(mainGameController.screenHeight/2)+40)
	traficLightController3 := traficLightController{
		lightsImage: lights3,
		lightsOptions: opl3,
		positionX: (mainGameController.screenWidth/2)-40,
		positionY: (mainGameController.screenHeight/2)+40,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController3)
	mainGameController.semaphores[2].options = opl3
	mainGameController.semaphores[2].image = lights3

	lights4, _, _ := ebitenutil.NewImageFromFile("roads/light.png", ebiten.FilterDefault)
	opl4 := &ebiten.DrawImageOptions{}
	opl4.GeoM.Translate(float64(mainGameController.screenWidth/2)-200, float64(mainGameController.screenHeight/2)-20)
	traficLightController4 := traficLightController{
		lightsImage: lights4,
		lightsOptions: opl4,
		positionX: (mainGameController.screenWidth/2)-200,
		positionY: (mainGameController.screenHeight/2)-20,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController4)
	mainGameController.semaphores[3].options = opl4
	mainGameController.semaphores[3].image = lights4
	
	lights5, _, _ := ebitenutil.NewImageFromFile("roads/light2.png", ebiten.FilterDefault)
	opl5 := &ebiten.DrawImageOptions{}
	opl5.GeoM.Translate(float64(mainGameController.screenWidth/2)+120, float64(mainGameController.screenHeight/2)-170)
	traficLightController5 := traficLightController{
		lightsImage: lights5,
		lightsOptions: opl5,
		positionX: (mainGameController.screenWidth/2)+120,
		positionY: (mainGameController.screenHeight/2)-170,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController5)
	mainGameController.semaphores[4].options = opl5
	mainGameController.semaphores[4].image = lights5
	
	lights6, _, _ := ebitenutil.NewImageFromFile("roads/light2.png", ebiten.FilterDefault)
	opl6 := &ebiten.DrawImageOptions{}
	opl6.GeoM.Translate(float64(mainGameController.screenWidth/2)+40, float64(mainGameController.screenHeight/2)-10)
	traficLightController6 := traficLightController{
		lightsImage: lights6,
		lightsOptions: opl6,
		positionX: (mainGameController.screenWidth/2)+40,
		positionY: (mainGameController.screenHeight/2)-10,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController6)
	mainGameController.semaphores[5].options = opl6
	mainGameController.semaphores[5].image = lights6

	lights7, _, _ := ebitenutil.NewImageFromFile("roads/light2.png", ebiten.FilterDefault)
	opl7 := &ebiten.DrawImageOptions{}
	opl7.GeoM.Translate(float64(mainGameController.screenWidth/2)-50, float64(mainGameController.screenHeight/2)-10)
	traficLightController7 := traficLightController{
		lightsImage: lights7,
		lightsOptions: opl7,
		positionX: (mainGameController.screenWidth/2)-50,
		positionY: (mainGameController.screenHeight/2)-10,
	}
	mainGameController.traficLightControllers = append(mainGameController.traficLightControllers, traficLightController7)
	mainGameController.semaphores[6].options = opl7
	mainGameController.semaphores[6].image = lights7
	
	//Light1 -40 -40
	//Light2 200 40
	//Light3 -40 40
	//Light4 -200 -40
	//Light5 40 -200
	//Light6 -40 -40
	//Light6 -40 -40
}

//Game Loo
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	//This part of the code draws the static road
	drawBoard(screen)

	//This part of the code draws all the cars in the screen and update the position of the cars
	for i := 0; i < mainGameController.numberOfCars; i++ {
		mainGameController.cars[i].options.GeoM.Translate(mainGameController.cars[i].speedX, mainGameController.cars[i].speedY)
		if mainGameController.cars[i].status == 1 {
			tempString := "Carro " + strconv.Itoa(i) + ": Ruta terminada"
			text.Draw(screen, tempString , normalFont, 15 , int(mainGameController.cars[i].positionY+10)*(i+1)*2, color.White)
		}else if mainGameController.cars[i].status == 0 {
			tempString := "Carro " + strconv.Itoa(i) + ": En ruta"
			text.Draw(screen, tempString, normalFont, 15 , int(mainGameController.cars[i].positionY+10)*(i+1)*2, color.White)
		}
		screen.DrawImage(mainGameController.cars[i].image, &mainGameController.cars[i].options)
	}

	//This part draws the text of the trafic lights
	for i := 0; i < mainGameController.numberOfSemaphores; i++ {
		if mainGameController.semaphores[i].color == 1 {
			text.Draw(screen, "Rojo", normalFont, mainGameController.traficLightControllers[i].positionX , mainGameController.traficLightControllers[i].positionY-7, color.White)
		}else if mainGameController.semaphores[i].color == 0 {
			text.Draw(screen, "Verde", normalFont, mainGameController.traficLightControllers[i].positionX , mainGameController.traficLightControllers[i].positionY-7, color.White)
		}	
		screen.DrawImage(mainGameController.semaphores[i].image,mainGameController.semaphores[i].options)
	}

	return nil
}

func main() {
	//Initialize seed for real random numbers
	rand.Seed(time.Now().UnixNano())
	//Status of the game
	//0 Game not started
	//1 Game started
	//2 Game finished
	gameStarted = 0
	//Initialization of game controller
	mainGameController = gameController{
		numberOfCars:       5,
		numberOfSemaphores: 4,
		screenWidth: 620,	
		screenHeight: 400,

	}
	//Initialization of starting points
	positionX1 := 0.0
	positionY1 := 280.0

	positionX2 := 280.0
	positionY2 := 360.0

	positionX3 := 520.0
	positionY3 := 360.0

	positionX4 := 520.0
	positionY4 := 0.0

	mainGameController.startingPositionsX = append(mainGameController.startingPositionsX,positionX1)
	mainGameController.startingPositionsY = append(mainGameController.startingPositionsY,positionY1)

	mainGameController.startingPositionsX = append(mainGameController.startingPositionsX,positionX2)
	mainGameController.startingPositionsY = append(mainGameController.startingPositionsY,positionY2)

	mainGameController.startingPositionsX = append(mainGameController.startingPositionsX,positionX3)
	mainGameController.startingPositionsY = append(mainGameController.startingPositionsY,positionY3)

	mainGameController.startingPositionsX = append(mainGameController.startingPositionsX,positionX4)
	mainGameController.startingPositionsY = append(mainGameController.startingPositionsY,positionY4)
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
			tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
			if err != nil {
				log.Fatal(err)
			}
			normalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    12,
				DPI:     dpi,
				Hinting: font.HintingFull,
			})
			for i := 0; i < mainGameController.numberOfCars; i++ {
				//Default values for cars
				randomPosition := (rand.Intn(4) + 0)
				currentImage, _, err := ebitenutil.NewImageFromFile("roads/car.png", ebiten.FilterDefault)
				if err != nil {
					log.Fatal(err)
				}
				tempCar := car{
					//Random speed from 1 to 10
					speedX: (rand.Float64() * (2.5-.8)),
					//speedY: (rand.Float64() * (2.5-.8)),
					speedY: 0,
					status: 0,
					start: randomPosition,
					rotationType: 0,
					//Image of the car
					image: currentImage,
					positionX: 0,
					positionY: 0,
					//Draw Options of the car
					options: ebiten.DrawImageOptions{},
				}
				mainGameController.cars = append(mainGameController.cars, tempCar)
				if randomPosition == 0{
					mainGameController.cars[i].options.GeoM.Rotate(float64((math.Pi / 360)*360))
					mainGameController.cars[i].rotationType = 0;
				}else if randomPosition == 1{
					mainGameController.cars[i].options.GeoM.Rotate(float64((math.Pi / 360)*180))
					mainGameController.cars[i].rotationType = 1;
				}else if randomPosition == 2{
					mainGameController.cars[i].rotationType = 2;
				}else if randomPosition == 3{
					mainGameController.cars[i].rotationType = 3;
					mainGameController.cars[i].options.GeoM.Rotate(float64((math.Pi / 360)*540))
				}
				mainGameController.cars[i].options.GeoM.Translate(mainGameController.startingPositionsX[randomPosition], mainGameController.startingPositionsY[randomPosition])
				go carBehavior(i)
			}
			for i := 0; i < 7; i++ {
				//Default values for semaphores
				tempSemaphore := semaphore{
					//Random color from 1 to 0
					color:   (rand.Intn(2) + 0),
					counter: 0,
				}
				mainGameController.semaphores = append(mainGameController.semaphores, tempSemaphore)
				go semaphoreBehavior(i)
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
