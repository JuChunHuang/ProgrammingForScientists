package main

import (
	"fmt"
	"testing"
)

func TestConstructQuadTree(t *testing.T) {
	u := CreateCustomUniverse()
	qt := ConstructQuadTree(u.stars, u.width)
	UpdateDummyStar(qt.root)
	fmt.Println(qt.root.star)
}

func TestUpdateVelocity(t *testing.T) {
	type test struct {
		s      Star
		time   float64
		answer OrderedPair
	}

	var s = Star{OrderedPair{100, 100}, OrderedPair{2, 4}, OrderedPair{1, 0}, 1, 1, 0, 0, 0}
	time := 1.0
	var ans = OrderedPair{3, 4}
	var test_case = test{s, time, ans}

	outcome := test_case.s.UpdateVelocity(test_case.time)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestUpdatePosition(t *testing.T) {
	type test struct {
		s      Star
		time   float64
		answer OrderedPair
	}

	var s = Star{OrderedPair{100, 100}, OrderedPair{2, 4}, OrderedPair{1, 0}, 1, 1, 0, 0, 0}
	time := 1.0
	var ans = OrderedPair{102.5, 104.0}
	var test_case = test{s, time, ans}

	outcome := test_case.s.UpdatePosition(test_case.time)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestInWhich(t *testing.T) {
	type test struct {
		p      OrderedPair
		n      Node
		answer int
	}

	var p = OrderedPair{76, 80}
	var n = Node{nil, nil, Quadrant{0, 100, 100}}
	ans := 3
	var test_case = test{p, n, ans}

	outcome := InWhich(test_case.p, test_case.n)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%d) but the answer is: (%d)", outcome, test_case.answer)
	} else {
		fmt.Println("Pass!")
	}
}

func TestCalculateCOM(t *testing.T) {
	type test struct {
		p1, p2 OrderedPair
		m1, m2 float64
		answer OrderedPair
	}
	var p1 = OrderedPair{1, 2}
	var p2 = OrderedPair{3, 4}
	m1 := 1.0
	m2 := 1.0
	ans := OrderedPair{2, 3}
	var test_case = test{p1, p2, m1, m2, ans}

	outcome := CalculateCOM(test_case.p1, test_case.p2, test_case.m1, test_case.m2)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func CreateCustomUniverse() *Universe {
	var A, B, C, D, E, F, G Star
	A.position.x, A.position.y = 1, 15
	B.position.x, B.position.y = 1, 13
	C.position.x, C.position.y = 6, 9
	D.position.x, D.position.y = 9, 9
	E.position.x, E.position.y = 11, 6
	F.position.x, F.position.y = 2, 5
	G.position.x, G.position.y = 9, 2

	A.mass = 1
	B.mass = 1
	C.mass = 1
	D.mass = 1
	E.mass = 1
	F.mass = 1
	G.mass = 1

	var custom_universe Universe
	custom_universe.width = 16
	custom_universe.AddStar(A)
	custom_universe.AddStar(B)
	custom_universe.AddStar(C)
	custom_universe.AddStar(D)
	custom_universe.AddStar(E)
	custom_universe.AddStar(F)
	custom_universe.AddStar(G)

	return &custom_universe
}

func (u *Universe) AddStar(s Star) {
	u.stars = append(u.stars, &s)
}
