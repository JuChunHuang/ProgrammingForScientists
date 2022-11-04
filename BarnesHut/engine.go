package main

// BarnesHut is our highest level function.
// Input: initial Universe object, a number of generations, and a time interval.
// Output: collection of Universe objects corresponding to updating the system over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, num_gens int, time, theta float64) []*Universe {
	time_points := make([]*Universe, num_gens+1)
	time_points[0] = initialUniverse

	for i := 1; i <= num_gens; i++ {
		time_points[i] = UpdateUniverse(time_points[i-1], time, theta)
	}

	return time_points
}

// UpdateUniverse updates a given Universe over a specified time interval (in seconds).
// Input: a Universe object, a float time, and a theta parameter.
// Output: a Universe object over time seconds.
func UpdateUniverse(current_universe *Universe, time, theta float64) *Universe {
	new_universe := current_universe.CopyUniverse()

	// constrruct quadtree
	qt := ConstructQuadTree(new_universe.stars, new_universe.width)
	// update the position and the mass of internal nodes (dummy stars)
	UpdateDummyStar(qt.root)

	//range over all stars in the universe and update their acceleration, velocity, and position
	for i := range new_universe.stars {
		new_universe.stars[i].acceleration = new_universe.stars[i].UpdateAcceleration(qt, theta)
		new_universe.stars[i].velocity = new_universe.stars[i].UpdateVelocity(time)
		new_universe.stars[i].position = new_universe.stars[i].UpdatePosition(time)
	}

	return new_universe
}

// UpdateVelocity updates the velocity of a given star object over a specified time interval (in seconds).
// Input: a star object and a time step (float64).
// Output: the orderedPair corresponding to the velocity of this object after a single time step, using the star's current acceleration.
func (s *Star) UpdateVelocity(time float64) OrderedPair {
	var new_velocity OrderedPair
	new_velocity.x = s.acceleration.x*time + s.velocity.x
	new_velocity.y = s.acceleration.y*time + s.velocity.y

	return new_velocity
}

// UpdatePosition updates the position of a given star object over a specified time interval (in seconds).
// Input: a star s and a time step (float64).
// Output: the OrderedPair corresponding to the updated position of the star after a single time step, using the star's current acceleration and velocity.
func (s *Star) UpdatePosition(time float64) OrderedPair {
	var new_position OrderedPair
	new_position.x = s.acceleration.x*time*time/2 + s.velocity.x*time + s.position.x
	new_position.y = s.acceleration.y*time*time/2 + s.velocity.y*time + s.position.y

	return new_position
}

// UpdateAcceleration updates star's acceleration over a specified time interval (in seconds).
// Input: a quadtree and a theta parameter.
// Output: the net acceleration on s due to net force calculated by stars in the quadtree.
func (s *Star) UpdateAcceleration(qt *QuadTree, theta float64) OrderedPair {
	var accel OrderedPair

	//compute net force vector acting on s
	force := s.ComputeNetForce(qt, theta)

	//now, calculate acceleration (F = ma)
	accel.x = force.x / s.mass
	accel.y = force.y / s.mass

	return accel
}

// ComputeNetForce sums the all forces based on the quadtree acting on the star s.
// Input: a quadtree and a theta parameter.
// Output: the net force vector (OrderedPair) acting on the given star.
func (s *Star) ComputeNetForce(qt *QuadTree, theta float64) OrderedPair {
	var net_force OrderedPair
	// use BFS traversal to examine each node
	queue := make([]*Node, 1)
	queue[0] = qt.root

	for len(queue) != 0 {
		cur := queue[0]
		if cur.children == nil && s.IsSameStar(cur.star) == false {
			// if the current node is a leaf node with a star
			F := s.ComputeForce(cur.star)
			net_force.AddNewForce(F)
		} else {
			// if the current node is an internal node
			param := s.CalculateTheta(cur)
			if param > theta {
				for i := range cur.children {
					if cur.children[i].star != nil {
						queue = append(queue, cur.children[i])
					}
				}
			} else {
				F := s.ComputeForce(cur.star)
				net_force.AddNewForce(F)
			}
		}
		queue = queue[1:]
	}

	return net_force
}

// ComputeForce computes the force acting on star s.
// Input: another star.
// Output: the force acting on star s.
func (s *Star) ComputeForce(new_star *Star) OrderedPair {
	var force OrderedPair

	d := Distance(s.position, new_star.position)
	F := G * s.mass * new_star.mass / (d * d)
	deltaX := new_star.position.x - s.position.x
	deltaY := new_star.position.y - s.position.y

	force.x = F * deltaX / d
	force.y = F * deltaY / d

	return force
}
