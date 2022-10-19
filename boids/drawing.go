package main

import (
	"canvas"
	"image"
)

func AnimateSystem(timePoints []Sky, canvasWidth, imageFrequency int) []image.Image {
	images := make([]image.Image, 0, len(timePoints))

	for i := range timePoints {
		// only draw if current index of sky
		if i%imageFrequency == 0 {
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}

	return images
}

func DrawToCanvas(s Sky, canvasWidth int) image.Image {
	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the boids and draw them.
	for _, b := range s.boids {
		c.SetFillColor(canvas.MakeColor(255, 255, 255))
		cx := (b.position.x / s.width) * float64(canvasWidth)
		cy := (b.position.y / s.width) * float64(canvasWidth)
		c.Circle(cx, cy, 5)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
