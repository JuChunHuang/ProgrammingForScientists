package main

import (
	"canvas"
	"image"
)

//AnimateSystem takes a slice of Universe objects along with a canvas width parameter and a frequency parameter.
//Every frequency steps, it generates a slice of images corresponding to drawing each Universe on a canvasWidth x canvasWidth canvas.
//A scaling factor is a final input that is used to scale the stars big enough to see them.
func AnimateSystem(time_points []*Universe, canvas_width, frequency int, scaling_factor float64) []image.Image {
	images := make([]image.Image, 0)

	if len(time_points) == 0 {
		panic("Error: no Universe objects present in AnimateSystem.")
	}

	// for every universe, draw to canvas and grab the image
	for i := range time_points {
		if i%frequency == 0 {
			if time_points[i] != nil {
				images = append(images, time_points[i].DrawToCanvas(canvas_width, scaling_factor))
			}
		}
	}

	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels.
//A scaling factor is needed to make the stars big enough to see them.
func (u *Universe) DrawToCanvas(canvas_width int, scaling_factor float64) image.Image {
	if u == nil {
		panic("Can't Draw a nil Universe.")
	}

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvas_width, canvas_width)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvas_width, canvas_width)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range u.stars {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		cx := (b.position.x / u.width) * float64(canvas_width)
		cy := (b.position.y / u.width) * float64(canvas_width)
		r := scaling_factor * (b.radius / u.width) * float64(canvas_width)
		c.Circle(cx, cy, r)
		c.Fill()
	}
	// we want to return an image!
	return c.GetImage()
}
