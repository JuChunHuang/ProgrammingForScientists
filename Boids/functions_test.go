package main

import (
	"fmt"
	"math"
	"testing"
)

func TestChangeDueToSeparation(t *testing.T) {
	type test struct {
		p1, p2            OrderedPair
		d                 float64
		separation_factor float64
		answer            OrderedPair
	}

	var p1 = OrderedPair{100, 100}
	var p2 = OrderedPair{40, 20}
	d := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	sep := 1.5
	var ans = OrderedPair{0.009, 0.012}
	var test_case = test{p1, p2, d, sep, ans}

	outcome := ChangeDueToSeparation(test_case.p1, test_case.p2, test_case.d, test_case.separation_factor)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestChangeDueToAlignment(t *testing.T) {
	type test struct {
		v                OrderedPair
		d                float64
		alignment_factor float64
		answer           OrderedPair
	}

	var v = OrderedPair{3, 4}
	d := 10.0
	ali := 1.0
	var ans = OrderedPair{0.3, 0.4}
	var test_case = test{v, d, ali, ans}

	outcome := ChangeDueToAlignment(test_case.v, test_case.d, test_case.alignment_factor)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestChangeDueToCohesion(t *testing.T) {
	type test struct {
		p1, p2            OrderedPair
		d                 float64
		separation_factor float64
		answer            OrderedPair
	}

	var p1 = OrderedPair{100, 100}
	var p2 = OrderedPair{40, 20}
	d := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	coh := 0.1
	var ans = OrderedPair{-0.06, -0.08}
	var test_case = test{p1, p2, d, coh, ans}

	outcome := ChangeDueToCohesion(test_case.p1, test_case.p2, test_case.d, test_case.separation_factor)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestUpdateVelocity(t *testing.T) {
	type test struct {
		b              Boid
		max_boid_speed float64
		time           float64
		answer         OrderedPair
	}

	var b = Boid{OrderedPair{100, 100}, OrderedPair{2, 4}, OrderedPair{1, 0}}
	maxSpeed := 2.5
	time := 1.0
	var ans = OrderedPair{1.5, 2.0}
	var test_case = test{b, maxSpeed, time, ans}

	outcome := UpdateVelocity(test_case.b, test_case.max_boid_speed, test_case.time)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestUpdatePosition(t *testing.T) {
	type test struct {
		b      Boid
		time   float64
		answer OrderedPair
	}

	var b = Boid{OrderedPair{100, 100}, OrderedPair{2, 4}, OrderedPair{1, 0}}
	time := 1.0
	var ans = OrderedPair{102.5, 104.0}
	var test_case = test{b, time, ans}

	outcome := UpdatePosition(test_case.b, test_case.time)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestUpdateTorusPosition(t *testing.T) {
	type test struct {
		b      Boid
		width  float64
		answer OrderedPair
	}

	var b = Boid{OrderedPair{-20, 2010}, OrderedPair{2, 4}, OrderedPair{1, 0}}
	width := 2000.0
	var ans = OrderedPair{1980, 10}
	var test_case = test{b, width, ans}

	outcome := UpdateTorusPosition(test_case.b, test_case.width)
	if outcome != test_case.answer {
		t.Errorf("Error! Output: (%f, %f) but the answer is: (%f, %f)", outcome.x, outcome.y, test_case.answer.x, test_case.answer.y)
	} else {
		fmt.Println("Pass!")
	}
}

func TestInBoard(t *testing.T) {
	type test struct {
		b      Boid
		width  float64
		answer bool
	}
	var b = Boid{OrderedPair{-20, 2010}, OrderedPair{2, 4}, OrderedPair{1, 0}}
	width := 2000.0
	ans := false
	var test_case = test{b, width, ans}

	outcome := InBoard(test_case.b, test_case.width)
	if outcome != test_case.answer {
		t.Errorf("Error! Wrong InBoard()")
	} else {
		fmt.Println("Pass!")
	}

}
