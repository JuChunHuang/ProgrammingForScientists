package main

import "math"

// CopyNode makes a deep copy of a node.
func (current_node *Node) CopyNode() *Node {
	children_copy := make([]*Node, 0)
	for i := range current_node.children {
		children_copy = append(children_copy, current_node.children[i])
	}
	star_copy := current_node.star.CopyStar()
	sector_copy := Quadrant{current_node.sector.x, current_node.sector.y, current_node.sector.width}
	n := Node{children_copy, star_copy, sector_copy}

	return &n
}

// CopyStar makes a deep copy of a star.
func (current_star *Star) CopyStar() *Star {
	var new_star Star

	new_star.position.x = current_star.position.x
	new_star.position.y = current_star.position.y
	new_star.velocity.x = current_star.velocity.x
	new_star.velocity.y = current_star.velocity.y
	new_star.acceleration.x = current_star.acceleration.x
	new_star.acceleration.y = current_star.acceleration.y
	new_star.mass = current_star.mass
	new_star.radius = current_star.radius
	new_star.red = current_star.red
	new_star.blue = current_star.blue
	new_star.green = current_star.green

	return &new_star
}

// CopyUniverse makes a deep copy of a Universe.
func (current_universe *Universe) CopyUniverse() *Universe {
	var new_universe Universe
	new_universe.width = current_universe.width
	new_universe.stars = make([]*Star, len(current_universe.stars))
	for i := range new_universe.stars {
		new_universe.stars[i] = current_universe.stars[i].CopyStar()
	}

	return &new_universe
}

// IsSameStar whether the given two stars are the same star or not.
func (s1 *Star) IsSameStar(s2 *Star) bool {
	if s1.position == s2.position && s1.velocity == s2.velocity && s1.acceleration == s2.acceleration && s1.mass == s2.mass && s1.radius == s2.radius && s1.red == s2.red && s1.blue == s2.blue && s1.green == s2.green {
		return true
	}

	return false
}

// Distance calculates the distance between two stars.
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// AddNewForce sums two forces.
func (total_force *OrderedPair) AddNewForce(new_force OrderedPair) {
	total_force.x += new_force.x
	total_force.y += new_force.y
}

// CalculateTheta computes theta parameter.
// Input: a node object.
// Output: the theta parameter.
func (s *Star) CalculateTheta(node *Node) float64 {
	d := Distance(s.position, node.star.position)

	return float64(node.sector.width) / d
}
