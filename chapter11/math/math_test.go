package math

import "testing"

type testpair struct {
	values[] float64
	expected float64
}

var tests_avg = []testpair{
	{[]float64{1,2}, 1.5},
	{[]float64{1,1,1,1,1}, 1},
	{[]float64{-1,1}, 0},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests_avg {
		v := Average(pair.values)
		if v != pair.expected {
			t.Errorf("For %v, expected %v, got %v", pair.values, pair.expected, v)
		}
	}
}

var tests_min = []testpair{
	{[]float64{0,0,0,0,0}, 0},
	{[]float64{5,4,1,2,3}, 1},
	{[]float64{-5,-4,-1,-2,-3}, -5},
	{[]float64{}, 0},
}

func TestMin(t *testing.T) {
	for _, pair := range tests_min {
		v := Min(pair.values)
		if v != pair.expected {
			t.Errorf("For %v, expected %v, got %v", pair.values, pair.expected, v)
		}
	}
}

var tests_max = []testpair {
	{[]float64{0,0,0,0,0}, 0},
	{[]float64{5,4,1,2,3}, 5},
	{[]float64{-5,-4,-1,-2,-3}, -1},
	{[]float64{}, 0},
}

func TestMax(t *testing.T) {
	for _, pair := range tests_max {
		v := Max(pair.values)
		if v != pair.expected {
			t.Errorf("For %v, expected %v, got %v", pair.values, pair.expected, v)
		}
	}
}