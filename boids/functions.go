package main

import (
	"math"
)

// SimulateBoids simulates the boids system over numGens generations starting with initialSky using a time step.
// Input: an initial Sky object, a number of generations, and a time interval (in seconds).
// Output: a slice of numGens + 1 total Sky objects.
func SimulateBoids(initial_sky Sky, num_gens int, time_step float64) []Sky {
	time_points := make([]Sky, num_gens+1)
	time_points[0] = initial_sky

	//now range over the number of generations and update the Sky each time
	for i := 1; i <= num_gens; i++ {
		time_points[i] = UpdateSky(time_points[i-1], time_step)
	}

	return time_points
}

// UpdateSkye updates a given Sky over a specified time interval (in seconds).
// Input: a Sky object and a float time.
// Output: a Sky object corresponding to simulating gravity over time seconds, assuming that acceleration is constant over this time.
func UpdateSky(current_sky Sky, time_step float64) Sky {
	new_sky := CopySky(current_sky)

	for i := range new_sky.boids {
		// range over all boids in the sky and update their acceleration, velocity, and position
		new_sky.boids[i].acceleration = UpdateAcceleration(current_sky, new_sky.boids[i])
		new_sky.boids[i].velocity = UpdateVelocity(new_sky.boids[i], new_sky.max_boid_speed, time_step)
		new_sky.boids[i].position = UpdatePosition(new_sky.boids[i], time_step)

		// examine whether boids fly off the edge of the board, if so, let them back
		if !InBoard(new_sky.boids[i], new_sky.width) {
			new_sky.boids[i].position = UpdateTorusPosition(new_sky.boids[i], new_sky.width)
		}
	}

	return new_sky
}

// UpdateAcceleration updates boid's acceleration over a specified time interval (in seconds).
// Input: Sky object and a boid b
// Output: the net acceleration on b due to net force calculated by every boid in the Sky
func UpdateAcceleration(current_sky Sky, b Boid) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on b
	force := ComputeNetForce(current_sky, b)

	//now, calculate acceleration. Since mass is equal to 1, F = a.
	accel.x = force.x
	accel.y = force.y

	return accel
}

// ComputeNetForce sums the all forces within a threshold distance acting on the boid b
// Input: A Sky objects and an individual boid
// Output: the net force vector (OrderedPair) acting on the given boid
func ComputeNetForce(current_sky Sky, b Boid) OrderedPair {
	var net_force OrderedPair
	num_under_thres := 0
	for i := range current_sky.boids {
		// only do a force computation if current boid is not the input boid b
		if current_sky.boids[i] != b {
			d := Distance(b.position, current_sky.boids[i].position)
			// only do a force computation if current boid stays within a threshold distance of b
			if d <= current_sky.proximity {
				num_under_thres += 1
				force := ComputeForce(b, current_sky.boids[i], d, current_sky.separation_factor, current_sky.alignment_factor, current_sky.cohesion_factor)

				//now add its components into net force components
				net_force.x += force.x
				net_force.y += force.y
			}
		}
	}

	// average all these forces to obtain a final force
	if num_under_thres > 0 {
		net_force.x /= float64(num_under_thres)
		net_force.y /= float64(num_under_thres)
	} else {
		return net_force
	}

	return net_force
}

// ComputeForce computes the total three forces
// Input: two Boid objects and three force factors
// Output: the force due to three rules (as a vector) acting on b1 subject to b2.
func ComputeForce(b1, b2 Boid, d, separation_factor, alignment_factor, cohesion_factor float64) OrderedPair {
	var force OrderedPair
	separation_force := ChangeDueToSeparation(b1.position, b2.position, d, separation_factor)
	alignment_force := ChangeDueToAlignment(b2.velocity, d, alignment_factor)
	cohesion_force := ChangeDueToCohesion(b1.position, b2.position, d, cohesion_factor)

	force.x = separation_force.x + alignment_force.x + cohesion_force.x
	force.y = separation_force.y + alignment_force.y + cohesion_force.y

	return force
}

// ChangeDueToSeparation calculates the force due to separation of nearby boids which are within a threshold distance
// Input: two positions of the two Boid objects, a threshold distance, and a factor that dictates the magnitude of the separation force
// Output: the force due to separation rule acting on p1 subject to p2
func ChangeDueToSeparation(p1, p2 OrderedPair, d, separationd_factor float64) OrderedPair {
	var separation_force OrderedPair
	separation_force.x = separationd_factor * (p1.x - p2.x) / (d * d)
	separation_force.y = separationd_factor * (p1.y - p2.y) / (d * d)

	return separation_force
}

// ChangeDueToAlignment calculates the force due to alignment of nearby boids which are within a threshold distance
// Input: one velocity vector of the nearby boid, a threshold distance, and a factor that dictates the magnitude of the alignment force
// Output: the force due to alignment rule due to the nearby boid
func ChangeDueToAlignment(v OrderedPair, d, alignment_factor float64) OrderedPair {
	var alignment_force OrderedPair
	alignment_force.x = alignment_factor * v.x / d
	alignment_force.y = alignment_factor * v.y / d

	return alignment_force
}

// ChangeDueToCohesion calculates the force due to cohesion of nearby boids which are within a threshold distance
// Input: two positions of the two Boid objects, a threshold distance, and a factor that dictates the magnitude of the cohesion force
// Output: the force due to cohesion rule acting on p1 subject to p2
func ChangeDueToCohesion(p1, p2 OrderedPair, d, cohesion_factor float64) OrderedPair {
	var cohesion_force OrderedPair
	cohesion_force.x = cohesion_factor * (p2.x - p1.x) / d
	cohesion_force.y = cohesion_factor * (p2.y - p1.y) / d

	return cohesion_force
}

// UpdateVelocity updates the velocity of a given Boid object over a specified time interval (in seconds) and ensure each boid's speed does not exceed the maximum speed.
// Input: a Boid object, a maximum speed of the system, and a time step (float64).
// Output: the orderedPair corresponding to the velocity of this object after a single time step, using the boid's current acceleration.
func UpdateVelocity(b Boid, max_boid_speed, time float64) OrderedPair {
	var new_velocity OrderedPair
	new_velocity.x = b.acceleration.x*time + b.velocity.x
	new_velocity.y = b.acceleration.y*time + b.velocity.y

	net_velocity := math.Sqrt(new_velocity.x*new_velocity.x + new_velocity.y*new_velocity.y)
	// make sure each boid's speed is at most equal to the maximum speed of the system
	if net_velocity > max_boid_speed {
		reduced_coefficient := max_boid_speed / net_velocity
		new_velocity.x *= reduced_coefficient
		new_velocity.y *= reduced_coefficient
	}

	return new_velocity
}

// UpdatePosition updates the position of a given Boid object over a specified time interval (in seconds).
// Input: a Boid b and a time step (float64).
// Output: the OrderedPair corresponding to the updated position of the boid after a single time step, using the boid's current acceleration and velocity.
func UpdatePosition(b Boid, time float64) OrderedPair {
	var new_position OrderedPair
	new_position.x = 0.5*b.acceleration.x*time*time + b.velocity.x*time + b.position.x
	new_position.y = 0.5*b.acceleration.y*time*time + b.velocity.y*time + b.position.y

	return new_position
}

// UpdateTorusPosition updats the position of a given Boid object which flys off the edge of the board
// Input: a Boid b and the width of the Sky
// Output: the OrderedPair corresponding to the boid that should be confined to the Sky
func UpdateTorusPosition(b Boid, width float64) OrderedPair {
	var new_position OrderedPair
	new_position.x = b.position.x
	new_position.y = b.position.y

	if b.position.x > width {
		new_position.x = b.position.x - width
	} else if b.position.x < 0 {
		new_position.x = b.position.x + width
	}
	if b.position.y > width {
		new_position.y = b.position.y - width
	} else if b.position.y < 0 {
		new_position.y = b.position.y + width
	}

	return new_position
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
func CopySky(current_sky Sky) Sky {
	var new_sky Sky

	new_sky.width = current_sky.width
	new_sky.max_boid_speed = current_sky.max_boid_speed
	new_sky.proximity = current_sky.proximity
	new_sky.separation_factor = current_sky.separation_factor
	new_sky.alignment_factor = current_sky.alignment_factor
	new_sky.cohesion_factor = current_sky.cohesion_factor

	// make the new sky's slice of Boid objects
	numBoids := len(current_sky.boids)
	new_sky.boids = make([]Boid, numBoids)

	// copy all of the boids' fields into our new boids
	for i := range current_sky.boids {
		new_sky.boids[i].position.x = current_sky.boids[i].position.x
		new_sky.boids[i].position.y = current_sky.boids[i].position.y
		new_sky.boids[i].velocity.x = current_sky.boids[i].velocity.x
		new_sky.boids[i].velocity.y = current_sky.boids[i].velocity.y
		new_sky.boids[i].acceleration.x = current_sky.boids[i].acceleration.x
		new_sky.boids[i].acceleration.y = current_sky.boids[i].acceleration.y
	}

	return new_sky
}

//Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
