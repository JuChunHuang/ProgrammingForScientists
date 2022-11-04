package main

// ConstructQuadTree creates a new quadtree based on a series of stars.
// Input: a slice of stars and the width of the Universe.
// Output: a quadtree.
func ConstructQuadTree(stars []*Star, width float64) *QuadTree {
	qt := InitializeQuadTree(width)
	qt.Insert(stars)

	return qt
}

// InitializeQuadTree initializes a quadtree that contains a root and four empty children.
// Input: the width of the Universe
// Output: a quadtree.
func InitializeQuadTree(w float64) *QuadTree {
	var t QuadTree
	var r Node
	var s Star
	r.star = &s
	r.sector = Quadrant{x: 0, y: w, width: w}
	t.root = &r
	t.root.Split()

	return &t
}

// Insert inserts stars one by one to the quadtree based on quadrant.
// Input: a initialized quadtree and a slice of stars
// Output: None.
func (qt *QuadTree) Insert(stars []*Star) {
	for i := range stars {
		cur := qt.root
		// determine which quadrant the current star is located in.
		quadrant := InWhich(stars[i].position, *cur)
		next := cur.children[quadrant]

		for next.star != nil {
			// the position that is going to be inserted already has another star
			if next.children == nil {
				// the position is already the leaf node --> needs to create new children
				next.Split()
				// the old position will become an internal node --> replace with a dummy star
				// without doing this will cause severe pointer issues
				old_star_quadrant := InWhich(next.star.position, *next)
				next.children[old_star_quadrant].star = next.star

				var dummy_node *Node = next.CopyNode()
				cur.children[quadrant] = dummy_node
				cur = dummy_node
			} else {
				// traverse down
				cur = next
			}
			quadrant = InWhich(stars[i].position, *cur)
			next = cur.children[quadrant]
		}
		next.star = stars[i]
	}
}

// UpdateDummyStar updates the positions and the masses of internal nodes, which previously have not been processed in the Insert() function.
// Input: the root of the tree.
// Output: the position and the mass of the internal node (dummy star).
func UpdateDummyStar(n *Node) (float64, float64, float64) {
	var x float64
	var y float64
	var dummy_mass float64
	if n.children != nil {
		// this node is an internal node
		n.star.DummyStarInitialize()
		for i := range n.children {
			// range over all its children to calculate the center of mass and the final mass
			if n.children[i].star != nil {
				if n.children[i].children == nil {
					// this node is already the internal node of the last layer
					n.star.position = CalculateCOM(n.star.position, n.children[i].star.position, n.star.mass, n.children[i].star.mass)
					n.star.mass += n.children[i].star.mass
				} else {
					// traverse down (similar to DFS)
					x, y, dummy_mass = UpdateDummyStar(n.children[i])
					n.star.position = CalculateCOM(n.star.position, OrderedPair{x, y}, n.star.mass, dummy_mass)
					n.star.mass += dummy_mass
				}
			}
		}
	}

	return n.star.position.x, n.star.position.y, n.star.mass
}

// DummyStarInitialize resets the position and the mass of the dummy star.
// Input: a dummy star.
// Output: None.
func (s *Star) DummyStarInitialize() {
	s.position.x = 0
	s.position.y = 0
	s.mass = 0
}

// CalculateCOM calculates the center of mass of two stars.
// Input: positions and masses of two stars.
// Output: a center of mass (OrderedPair).
func CalculateCOM(p1, p2 OrderedPair, m1, m2 float64) OrderedPair {
	var COM_position OrderedPair
	mass_sum := m1 + m2
	COM_position.x = (p1.x*m1 + p2.x*m2) / mass_sum
	COM_position.y = (p1.y*m1 + p2.y*m2) / mass_sum

	return COM_position
}

// Split creates four empty children of the node and assigns the sectors.
// Input: a node.
// Output: None
func (n *Node) Split() {
	if n.children != nil {
		return
	}

	new_width := n.sector.width * 0.5
	var n1 Node
	var n2 Node
	var n3 Node
	var n4 Node
	n1.sector = Quadrant{x: n.sector.x, y: n.sector.y - new_width, width: new_width}             // NE
	n2.sector = Quadrant{x: n.sector.x + new_width, y: n.sector.y - new_width, width: new_width} // NW
	n3.sector = Quadrant{x: n.sector.x, y: n.sector.y, width: new_width}                         // SE
	n4.sector = Quadrant{x: n.sector.x + new_width, y: n.sector.y, width: new_width}             // SW
	n.children = append(n.children, &n1, &n2, &n3, &n4)
}

// InWhich determines the quadrant of the star based on its position in the Universe.
// Input: a position of a star and its parent node.
// Output: a int represents the quadrant.
func InWhich(p OrderedPair, n Node) int {
	half_width := n.sector.width * 0.5

	if p.y < n.sector.y-half_width {
		if p.x < n.sector.x+half_width {
			return 0 // NE
		} else {
			return 1 // NW
		}
	} else {
		if p.x < n.sector.x+half_width {
			return 2 // SE
		} else {
			return 3 // SW
		}
	}
}

// InField makes sure the star is inside the Universe.
// Input: a star and a root node.
// Output: boolean that represents the star is inside or outside the Universe.
func (n *Node) InField(s *Star) bool {
	if s.position.x >= n.sector.x && s.position.x <= n.sector.x+n.sector.width &&
		s.position.y >= n.sector.y-n.sector.width && s.position.y <= n.sector.y {
		return true
	}

	return false
}