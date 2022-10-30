package main

import (
	"canvas"
	"image"
)

func AnimateSystem(time_points []Sky, canvas_width, image_frequency int) []image.Image {
	images := make([]image.Image, 0, len(time_points))

	for i := range time_points {
		// only draw if current index of sky
		if i%image_frequency == 0 {
			images = append(images, DrawToCanvas(time_points[i], canvas_width))
		}
	}

	return images
}

func DrawToCanvas(s Sky, canvas_width int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvas_width, canvas_width)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvas_width, canvas_width)
	c.Fill()

	// range over all the boids and draw them.
	for _, b := range s.boids {
		c.SetFillColor(canvas.MakeColor(255, 255, 255))
		cx := (b.position.x / s.width) * float64(canvas_width)
		cy := (b.position.y / s.width) * float64(canvas_width)
		c.Circle(cx, cy, 5)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
