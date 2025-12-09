package main

import (
	"math"
	"slices"
	"strings"

	"github.com/m1thrandir225/aoc_2025/util"
)

type Point struct {
	x, y, z int
}

func EuclidianDistance(p1, p2 Point) float64 {
	x := math.Pow(float64(p1.x)-float64(p2.x), 2)
	y := math.Pow(float64(p1.y)-float64(p2.y), 2)
	z := math.Pow(float64(p1.z)-float64(p2.z), 2)

	return math.Sqrt(x + y + z)
}

func part1(input string) any {
	points := make([]Point, 0, len(input))

	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, ",")

		intSplits := util.StrToIntArr(splits)
		poit := Point{
			x: intSplits[0],
			y: intSplits[1],
			z: intSplits[2],
		}
		points = append(points, poit)
	}

	pointPairs := [][2]Point{}
	pointDistances := map[[2]Point]float64{}

	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			pair := [2]Point{points[i], points[j]}
			pointDistances[pair] = EuclidianDistance(points[i], points[j])
			pointPairs = append(pointPairs, pair)
		}
	}

	slices.SortFunc(pointPairs, func(a, b [2]Point) int {
		return int(EuclidianDistance(a[0], a[1])) - int(EuclidianDistance(b[0], b[1]))
	})

	pointConnections := make(map[Point][]Point)

	for i := range 1000 {
		pair := pointPairs[i]
		pointConnections[pair[0]] = append(pointConnections[pair[0]], pair[1])
		pointConnections[pair[1]] = append(pointConnections[pair[1]], pair[0])
	}

	g := groups(points, pointConnections)
	slices.Sort(g)
	slices.Reverse(g)

	res := 1
	for i := range g[:3] {
		res *= g[i]
	}
	return res
}

func groups(poss []Point, connections map[Point][]Point) []int {
	posSet := util.SetFrom(poss)
	var groupSizes []int
	seen := util.Set[Point]{}

	for len(posSet) > 0 {
		groupSize := 0
		q := util.NewQueue[Point]()
		q.Push(posSet.Slice()[0])
		for pos := range q.Seq {
			if seen.Contains(pos) {
				continue
			}
			seen.Add(pos)

			groupSize++
			posSet.Remove(pos)
			for _, next := range connections[pos] {
				q.Push(next)
			}
		}
		groupSizes = append(groupSizes, groupSize)
	}
	return groupSizes
}

func part2(input string) any {
	points := make([]Point, 0, len(input))

	for _, line := range strings.Split(input, "\n") {
		splits := strings.Split(line, ",")

		intSplits := util.StrToIntArr(splits)
		poit := Point{
			x: intSplits[0],
			y: intSplits[1],
			z: intSplits[2],
		}
		points = append(points, poit)
	}

	pointPairs := [][2]Point{}
	pointDistances := map[[2]Point]float64{}

	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			pair := [2]Point{points[i], points[j]}
			pointDistances[pair] = EuclidianDistance(points[i], points[j])
			pointPairs = append(pointPairs, pair)
		}
	}

	slices.SortFunc(pointPairs, func(a, b [2]Point) int {
		return int(EuclidianDistance(a[0], a[1])) - int(EuclidianDistance(b[0], b[1]))
	})

	pointConnections := make(map[Point][]Point)

	for _, pair := range pointPairs {
		pointConnections[pair[0]] = append(pointConnections[pair[0]], pair[1])
		pointConnections[pair[1]] = append(pointConnections[pair[1]], pair[0])

		if len(groups(points, pointConnections)) == 1 {
			return pair[0].x * pair[1].x
		}
	}
	return nil
}
