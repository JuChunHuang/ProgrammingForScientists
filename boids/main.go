package main

import (
	"fmt"
	"gifhelper"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// os.Args[1] is the number of boids in a sky
	numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	if numBoids < 0 {
		panic("Negative number of boids given.")
	}

	// os.Args[2] is the width of a sky
	skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		panic(err2)
	}

	// os.Args[3] is the initial speed of boids
	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3 != nil {
		panic(err3)
	}

	// os.Args[4] is maximum speed that a boid cannot exceed
	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4 != nil {
		panic(err4)
	}

	// os.Args[5] is the total number of generations
	numGens, err5 := strconv.Atoi(os.Args[5])
	if err5 != nil {
		panic(err5)
	}
	if numGens < 0 {
		panic("Negative number of generations given.")
	}

	// os.Args[6] is the threshold of distance of force
	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6 != nil {
		panic(err6)
	}

	// os.Args[7] is the magnitude factor due to separation rule
	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7 != nil {
		panic(err7)
	}

	// os.Args[8] is the magnitude factor due to alignment rule
	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8 != nil {
		panic(err8)
	}

	// os.Args[9] is the magnitude factor due to cohesion rule
	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
	if err9 != nil {
		panic(err9)
	}

	// os.Args[10] is the time step parameter
	timeStep, err10 := strconv.ParseFloat(os.Args[10], 64)
	if err10 != nil {
		panic(err10)
	}

	// os.Args[11] is the canvas width
	canvasWidth, err11 := strconv.Atoi(os.Args[11])
	if err11 != nil {
		panic(err11)
	}

	// os.Args[12] is how often to make a canvas
	imageFrequency, err12 := strconv.Atoi(os.Args[12])
	if err12 != nil {
		panic(err12)
	}

	fmt.Println("Command line arguments read successfully.")

	// declaring Sky and setting its fields.
	var initialSky Sky
	initialSky.width = skyWidth
	initialSky.boids = make([]Boid, numBoids)
	initialSky.maxBoidSpeed = maxBoidSpeed
	initialSky.proximity = proximity
	initialSky.separationFactor = separationFactor
	initialSky.alignmentFactor = alignmentFactor
	initialSky.cohesionFactor = cohesionFactor

	for i := range initialSky.boids {
		initialSky.boids[i].position.x = rand.Float64() * initialSky.width
		initialSky.boids[i].position.y = rand.Float64() * initialSky.width
		initialSky.boids[i].velocity.x = rand.Float64() * initialSpeed
		initialSky.boids[i].velocity.y = math.Sqrt(math.Pow(initialSpeed, 2) - math.Pow(initialSky.boids[i].velocity.x, 2))
		if rand.Intn(2) < 1 {
			initialSky.boids[i].velocity.x = -initialSky.boids[i].velocity.x
		}
		if rand.Intn(2) < 1 {
			initialSky.boids[i].velocity.y = -initialSky.boids[i].velocity.y
		}
	}

	fmt.Println("Simulating system.")

	timePoints := SimulateBoids(initialSky, numGens, timeStep)

	fmt.Println("Boids have been simulated!")

	fmt.Println("Ready to draw images.")

	images := AnimateSystem(timePoints, canvasWidth, imageFrequency)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF.")

	gifhelper.ImagesToGIF(images, "Boids")

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")
}
