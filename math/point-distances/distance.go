package main

import (
	"fmt"
	"math"
	"sort"
)

// given a rider location [x, y], find out 3 nearest drivers online

// src [x,y] [2,3]

// dest [x1,y1] [x2,y2] ... [xn, yn]

//[5,6] [1,3] ..

//PtA - Pt B =>

type Point struct {
	X, Y float64
}

func distance(p1, p2 Point) float64 {

	return math.Sqrt(math.Pow(math.Abs(p1.X-p2.X), 2) + math.Pow(math.Abs(p1.Y-p2.Y), 2))

}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	a := []Point{{5, 6}, {7, 8}, {1, 3}, {4, 5}}

	src := Point{2, 3}

	p := findKClosestPoints(a, src, 3)

	fmt.Println(p)

}

type DistancePoint struct {
	P        Point
	Distance float64
}

func findKClosestPoints(a []Point, src Point, k int) []Point {

	distances := make([]DistancePoint, len(a))

	for i, p := range a {
		distances[i].Distance = distance(p, src)
		distances[i].P = p
	}

	sort.SliceStable(distances, func(i, j int) bool { return distances[i].Distance < distances[j].Distance })

	var p []Point
	for i := 0; i < k; i++ {
		p = append(p, distances[i].P)
	}
	return p
}
