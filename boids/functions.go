package main

import (
	"math"
)

// SimulateBoids simulates the boids system over numGens generations starting with initialSky using a time step.
// Input: an initial Sky object, a number of generations, and a time interval (in seconds).
// Output: a slice of numGens + 1 total Sky objects.
func SimulateBoids(initialSky Sky, numGens int, timeStep float64) []Sky {
	timePoints := make([]Sky, numGens+1)
	timePoints[0] = initialSky

	//now range over the number of generations and update the Sky each time
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateSky(timePoints[i-1], timeStep)
	}

	return timePoints
}

// UpdateSkye updates a given Sky over a specified time interval (in seconds).
// Input: a Sky object and a float time.
// Output: a Sky object corresponding to simulating gravity over time seconds, assuming that acceleration is constant over this time.
func UpdateSky(currentSky Sky, timeStep float64) Sky {
	newSky := CopySky(currentSky)

	for i := range newSky.boids {
		// range over all boids in the sky and update their acceleration, velocity, and position
		newSky.boids[i].acceleration = UpdateAcceleration(currentSky, newSky.boids[i])
		newSky.boids[i].velocity = UpdateVelocity(newSky.boids[i], newSky.maxBoidSpeed, timeStep)
		newSky.boids[i].position = UpdatePosition(newSky.boids[i], timeStep)

		// examine whether boids fly off the edge of the board, if so, let them back
		if !InBoard(newSky.boids[i], newSky.width) {
			newSky.boids[i].position = UpdateTorusPosition(newSky.boids[i], newSky.width)
		}
	}

	return newSky
}

// UpdateAcceleration updates boid's acceleration over a specified time interval (in seconds).
// Input: Sky object and a boid b
// Output: the net acceleration on b due to net force calculated by every boid in the Sky
func UpdateAcceleration(currentSky Sky, b Boid) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeNetForce(currentSky, b)

	//now, calculate acceleration. Since mass is equal to 1, F = a.
	accel.x = force.x
	accel.y = force.y

	return accel
}

// ComputeNetForce sums the all forces within a threshold distance acting on the boid b
// Input: A Sky objects and an individual boid
// Output: the net force vector (OrderedPair) acting on the given boid
func ComputeNetForce(currentSky Sky, b Boid) OrderedPair {
	var netForce OrderedPair
	numUnderThres := 0
	for i := range currentSky.boids {
		// only do a force computation if current boid is not the input boid b
		if currentSky.boids[i] != b {
			d := Distance(b.position, currentSky.boids[i].position)
			// only do a force computation if current boid stays within a threshold distance of b
			if d <= currentSky.proximity {
				numUnderThres += 1
				force := ComputeForce(b, currentSky.boids[i], d, currentSky.separationFactor, currentSky.alignmentFactor, currentSky.cohesionFactor)

				//now add its components into net force components
				netForce.x += force.x
				netForce.y += force.y
			}
		}
	}

	// average all these forces to obtain a final force
	if numUnderThres > 0 {
		netForce.x /= float64(numUnderThres)
		netForce.y /= float64(numUnderThres)
	} else {
		return netForce
	}

	return netForce
}

// ComputeForce computes the total three forces
// Input: two Boid objects and three force factors
// Output: the force due to three rules (as a vector) acting on b1 subject to b2.
func ComputeForce(b1, b2 Boid, d, separationFactor, alignmentFactor, cohesionFactor float64) OrderedPair {
	var force OrderedPair
	separationForce := ChangeDueToSeparation(b1.position, b2.position, d, separationFactor)
	alignmentForce := ChangeDueToAlignment(b2.velocity, d, alignmentFactor)
	cohesionForce := ChangeDueToCohesion(b1.position, b2.position, d, cohesionFactor)

	force.x = separationForce.x + alignmentForce.x + cohesionForce.x
	force.y = separationForce.y + alignmentForce.y + cohesionForce.y

	return force
}

// ChangeDueToSeparation calculates the force due to separation of nearby boids which are within a threshold distance
// Input: two positions of the two Boid objects, a threshold distance, and a factor that dictates the magnitude of the separation force
// Output: the force due to separation rule acting on p1 subject to p2
func ChangeDueToSeparation(p1, p2 OrderedPair, d, separationFactor float64) OrderedPair {
	var separationForce OrderedPair
	separationForce.x = separationFactor * (p1.x - p2.x) / (d * d)
	separationForce.y = separationFactor * (p1.y - p2.y) / (d * d)

	return separationForce
}

// ChangeDueToAlignment calculates the force due to alignment of nearby boids which are within a threshold distance
// Input: one velocity vector of the nearby boid, a threshold distance, and a factor that dictates the magnitude of the alignment force
// Output: the force due to alignment rule due to the nearby boid
func ChangeDueToAlignment(v OrderedPair, d, alignmentFactor float64) OrderedPair {
	var alignmentForce OrderedPair
	alignmentForce.x = alignmentFactor * v.x / d
	alignmentForce.y = alignmentFactor * v.y / d

	return alignmentForce
}

// ChangeDueToCohesion calculates the force due to cohesion of nearby boids which are within a threshold distance
// Input: two positions of the two Boid objects, a threshold distance, and a factor that dictates the magnitude of the cohesion force
// Output: the force due to cohesion rule acting on p1 subject to p2
func ChangeDueToCohesion(p1, p2 OrderedPair, d, cohesionFactor float64) OrderedPair {
	var cohesionForce OrderedPair
	cohesionForce.x = cohesionFactor * (p2.x - p1.x) / d
	cohesionForce.y = cohesionFactor * (p2.y - p1.y) / d

	return cohesionForce
}

// UpdateVelocity updates the velocity of a given Boid object over a specified time interval (in seconds) and ensure each boid's speed does not exceed the maximum speed.
// Input: a Boid object, a maximum speed of the system, and a time step (float64).
// Output: the orderedPair corresponding to the velocity of this object after a single time step, using the boid's current acceleration.
func UpdateVelocity(b Boid, maxBoidSpeed, time float64) OrderedPair {
	var newVelocity OrderedPair
	newVelocity.x = b.acceleration.x*time + b.velocity.x
	newVelocity.y = b.acceleration.y*time + b.velocity.y

	netVelocity := math.Sqrt(newVelocity.x*newVelocity.x + newVelocity.y*newVelocity.y)
	// make sure each boid's speed is at most equal to the maximum speed of the system
	if netVelocity > maxBoidSpeed {
		reducedCoefficient := maxBoidSpeed / netVelocity
		newVelocity.x *= reducedCoefficient
		newVelocity.y *= reducedCoefficient
	}

	return newVelocity
}

// UpdatePosition updates the position of a given Boid object over a specified time interval (in seconds).
// Input: a Boid b and a time step (float64).
// Output: the OrderedPair corresponding to the updated position of the boid after a single time step, using the boid's current acceleration and velocity.
func UpdatePosition(b Boid, time float64) OrderedPair {
	var newPosition OrderedPair
	newPosition.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	newPosition.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y

	return newPosition
}

// UpdateTorusPosition updats the position of a given Boid object which flys off the edge of the board
// Input: a Boid b and the width of the Sky
// Output: the OrderedPair corresponding to the boid that should be confined to the Sky
func UpdateTorusPosition(b Boid, width float64) OrderedPair {
	var newPosition OrderedPair
	newPosition.x = b.position.x
	newPosition.y = b.position.y

	if b.position.x > width {
		newPosition.x = b.position.x - width
	} else if b.position.x < 0 {
		newPosition.x = b.position.x + width
	}
	if b.position.y > width {
		newPosition.y = b.position.y - width
	} else if b.position.y < 0 {
		newPosition.y = b.position.y + width
	}

	return newPosition
}

// InBoard examines whether the boid is flying away (outside the Sky)
func InBoard(b Boid, width float64) bool {
	if b.position.x > width || b.position.x < 0 || b.position.y > width || b.position.y < 0 {
		return false
	}

	return true
}

// CopySky
// Input: a Sky object
// Output: a new Sky object, all of whose fields are copied over into the new Sky's fields. (Deep copy)
func CopySky(currentSky Sky) Sky {
	var newSky Sky

	newSky.width = currentSky.width
	newSky.maxBoidSpeed = currentSky.maxBoidSpeed
	newSky.proximity = currentSky.proximity
	newSky.separationFactor = currentSky.separationFactor
	newSky.alignmentFactor = currentSky.alignmentFactor
	newSky.cohesionFactor = currentSky.cohesionFactor

	// make the new sky's slice of Boid objects
	numBoids := len(currentSky.boids)
	newSky.boids = make([]Boid, numBoids)

	// copy all of the boids' fields into our new boids
	for i := range currentSky.boids {
		newSky.boids[i].position.x = currentSky.boids[i].position.x
		newSky.boids[i].position.y = currentSky.boids[i].position.y
		newSky.boids[i].velocity.x = currentSky.boids[i].velocity.x
		newSky.boids[i].velocity.y = currentSky.boids[i].velocity.y
		newSky.boids[i].acceleration.x = currentSky.boids[i].acceleration.x
		newSky.boids[i].acceleration.y = currentSky.boids[i].acceleration.y
	}

	return newSky
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
