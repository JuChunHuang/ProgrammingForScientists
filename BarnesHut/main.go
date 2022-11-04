package main

import (
	"fmt"
	"gifhelper"
	"math"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	mode := os.Args[1]
	if mode == "galaxy" {
		GalaxySimulation()
	} else if mode == "jupiter" {
		JupiterSimulation()
	} else {
		CollisionSimulation()
	}
}

func JupiterSimulation() {
	var jupiter, io, europa, ganymede, callisto Star

	jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
	io.red, io.green, io.blue = 249, 249, 165
	europa.red, europa.green, europa.blue = 132, 83, 52
	ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
	callisto.red, callisto.green, callisto.blue = 0, 153, 76

	jupiter.mass = 1.898 * math.Pow(10, 27)
	io.mass = 8.9319 * math.Pow(10, 22)
	europa.mass = 4.7998 * math.Pow(10, 22)
	ganymede.mass = 1.4819 * math.Pow(10, 23)
	callisto.mass = 1.0759 * math.Pow(10, 23)

	jupiter.radius = 71000000
	io.radius = 1821000
	europa.radius = 1569000
	ganymede.radius = 2631000
	callisto.radius = 2410000

	jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
	io.position.x, io.position.y = 2000000000-421600000, 2000000000
	europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
	ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
	callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

	jupiter.velocity.x, jupiter.velocity.y = 0, 0
	io.velocity.x, io.velocity.y = 0, -17320
	europa.velocity.x, europa.velocity.y = -13740, 0
	ganymede.velocity.x, ganymede.velocity.y = 0, 10870
	callisto.velocity.x, callisto.velocity.y = 8200, 0

	// declaring universe and setting its fields.
	var jupiter_system Universe
	jupiter_system.width = 4000000000
	jupiter_system.stars = make([]*Star, 5)
	jupiter_system.stars[0] = (&jupiter).CopyStar()
	jupiter_system.stars[1] = (&io).CopyStar()
	jupiter_system.stars[2] = (&europa).CopyStar()
	jupiter_system.stars[3] = (&ganymede).CopyStar()
	jupiter_system.stars[4] = (&callisto).CopyStar()

	var num_gens int = 1000000
	var time float64 = 1.0
	var canvas_width int = 500
	var drawing_frequency int = 1000
	var theta float64 = 0.5
	var scaling_factor float64 = 5

	fmt.Println("Command line arguments read successfully.")

	fmt.Println("Simulating system.")

	time_points := BarnesHut(&jupiter_system, num_gens, time, theta)

	fmt.Println("Gravity has been simulated!")
	fmt.Println("Ready to draw images.")

	images := AnimateSystem(time_points, canvas_width, drawing_frequency, scaling_factor)

	fmt.Println("Images drawn!")

	fmt.Println("Making GIF.")

	gifhelper.ImagesToGIF(images, "jupiter")

	fmt.Println("Animated GIF produced!")

	fmt.Println("Exiting normally.")
}

func GalaxySimulation() {
	g0 := InitializeGalaxy(500, 4e21, 5e22, 4e22)
	width := 1.0e23
	galaxies := []Galaxy{g0}

	initial_universe := InitializeUniverse(galaxies, width)

	var num_gens int = 50000
	var time float64 = 2e14
	var theta float64 = 0.5
	var canvas_width int = 1000
	var drawing_frequency int = 1000
	var scaling_factor float64 = 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse

	time_points := BarnesHut(initial_universe, num_gens, time, theta)

	fmt.Println("Simulation run. Now drawing images.")
	image_list := AnimateSystem(time_points, canvas_width, drawing_frequency, scaling_factor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(image_list, "galaxy")
	fmt.Println("GIF drawn.")
}

func CollisionSimulation() {
	g0 := InitializeGalaxy(500, 4e21, 5e22, 4e22)
	g1 := InitializeGalaxy(500, 4e21, 4e22, 4e22)

	Push(&g0, OrderedPair{-100, 200})
	Push(&g1, OrderedPair{200, -100})

	width := 1.0e23
	galaxies := []Galaxy{g0, g1}

	initial_universe := InitializeUniverse(galaxies, width)

	var num_gens int = 12000
	var time float64 = 2e15
	var theta float64 = 0.5
	var canvas_width int = 800
	var draw_frequency int = 300
	var scaling_factor float64 = 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse

	time_points := BarnesHut(initial_universe, num_gens, time, theta)

	fmt.Println("Simulation run. Now drawing images.")
	image_list := AnimateSystem(time_points, canvas_width, draw_frequency, scaling_factor)

	fmt.Println("Images drawn. Now generating GIF.")
	gifhelper.ImagesToGIF(image_list, "collision")
	fmt.Println("GIF drawn.")
}
