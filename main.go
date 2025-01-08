package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Point struct {
	x, y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("%.2f %.2f", p.x, p.y)
}

func main() {
	var src *bufio.Reader
	if len(os.Args) > 1 {
		fn := os.Args[1]
		f, err := os.Open(fn)
		if err != nil {
			log.Fatalln(fmt.Errorf("error opening file: %w", err))
		}
		src = bufio.NewReader(f)
	} else {
		src = bufio.NewReader(os.Stdin)
	}

	for {
		var pointsNum int
		_, _ = fmt.Fscanf(src, "%d\n", &pointsNum)

		points := make([]*Point, 0, pointsNum)

		// Only way to exit loop.
		if pointsNum == 0 {
			return
		}

		// Clear points.
		points = points[0:0]

		for i := 0; i < pointsNum; i++ {
			var x, y float64
			_, _ = fmt.Fscanf(src, "%f %f\n", &x, &y)

			points = append(points, &Point{x, y})
		}

		// Initial sort by X.
		sort.Slice(points, func(i, j int) bool {
			return points[i].x < points[j].x
		})
		a, b, _ := findMinDist(points)
		fmt.Println(a, b)
	}
}

func findMinDist(localPoints []*Point) (*Point, *Point, float64) {
	localPointsLength := len(localPoints)
	if localPointsLength == 2 {
		p1 := localPoints[0]
		p2 := localPoints[1]
		return p1, p2, calcSquareDist(p1, p2)
	} else if localPointsLength == 3 {
		p1 := localPoints[0]
		p2 := localPoints[1]
		p3 := localPoints[2]

		dist1 := calcSquareDist(p1, p2)
		dist2 := calcSquareDist(p2, p3)
		dist3 := calcSquareDist(p3, p1)

		if dist1 < dist2 && dist1 < dist3 {
			return p1, p2, dist1
		} else if dist2 < dist3 {
			return p2, p3, dist2
		} else {
			return p3, p1, dist3
		}
	}

	median := localPointsLength / 2
	medianLocalPoint := localPoints[median]
	medianXVal := medianLocalPoint.x
	// Left Half.
	p1, p2, squareDist1 := findMinDist(localPoints[:median])
	// Right Half, middle point contained here.
	p3, p4, squareDist2 := findMinDist(localPoints[median:])

	var p1Ans *Point
	var p2Ans *Point
	var squareDistanceAns float64
	var delta float64
	if squareDist1 < squareDist2 {
		p1Ans = p1
		p2Ans = p2
		squareDistanceAns = squareDist1
		delta = math.Sqrt(squareDist1)
	} else {
		p1Ans = p3
		p2Ans = p4
		squareDistanceAns = squareDist2
		delta = math.Sqrt(squareDist2)
	}

	// ================ Middle Slab ================
	// Define middle slab bounds.
	leftBound := medianXVal - delta
	rightBound := medianXVal + delta

	// Search backwards through the local points,
	// add points to leftSlabPoints if within leftBound.
	leftBoundIndex := median
	for i := median - 1; i >= 0 && localPoints[i].x > leftBound; i-- {
		leftBoundIndex--
	}
	leftSlabPoints := localPoints[leftBoundIndex:median]

	// Search forward through the local points,
	// add points to rightSlabPoints if within rightBound.
	rightBoundIndex := median
	for i := rightBoundIndex; i < localPointsLength && localPoints[i].x < rightBound; i++ {
		rightBoundIndex++
	}
	rightSlabPoints := localPoints[median:rightBoundIndex]

	// Sort both slabs by Y.
	sort.Slice(leftSlabPoints, func(i, j int) bool {
		return leftSlabPoints[i].y < leftSlabPoints[j].y
	})
	sort.Slice(rightSlabPoints, func(i, j int) bool {
		return rightSlabPoints[i].y < rightSlabPoints[j].y
	})

	testPointDistance := func(left, right *Point) {
		if right.x-left.x < delta {
			sqDist := calcSquareDist(left, right)
			if sqDist < squareDistanceAns {
				delta = math.Sqrt(sqDist)
				squareDistanceAns = sqDist
				p1Ans = left
				p2Ans = right
			}
		}
	}
	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(leftSlabPoints) && rightIndex < len(rightSlabPoints) {
		leftPoint := leftSlabPoints[leftIndex]
		rightPoint := rightSlabPoints[rightIndex]
		if leftPoint.y <= rightPoint.y {
			// Right point is above left point.
			if rightPoint.y-leftPoint.y < delta {
				// Don't test distance if vertical distance is greater than delta.
				testPointDistance(leftPoint, rightPoint)
			}
			if nextRight := rightIndex + 1; nextRight < len(rightSlabPoints) {
				testPointDistance(leftPoint, rightSlabPoints[nextRight])
			}
			leftIndex++
		} else {
			// Left point is above right point.
			if leftPoint.y-rightPoint.y < delta {
				// Don't test distance if vertical distance is greater than delta.
				testPointDistance(leftPoint, rightPoint)
			}
			if nextLeft := leftIndex + 1; nextLeft < len(leftSlabPoints) {
				testPointDistance(leftSlabPoints[nextLeft], rightPoint)
			}
			rightIndex++
		}
	}

	// Sorting by Y breaks the sort by X, so we reorder it here.
	sort.Slice(leftSlabPoints, func(i, j int) bool {
		return leftSlabPoints[i].x < leftSlabPoints[j].x
	})
	sort.Slice(rightSlabPoints, func(i, j int) bool {
		return rightSlabPoints[i].x < rightSlabPoints[j].x
	})

	return p1Ans, p2Ans, squareDistanceAns
}

// calcSquareDist Returns the squared distance between two points.
// To get the actual distance, take the square root of this value.
// In many parts of the algorithm, calculating the square root is not necessary.
// This is a small optimization for that reason.
func calcSquareDist(a, b *Point) float64 {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
}
