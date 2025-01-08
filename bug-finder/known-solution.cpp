#include <iostream>
#include <vector>
#include <algorithm>
#include <cmath>
#include <limits>

struct Point {
	double x, y;
	int id; // To track original points for output
};

// Compare points by x-coordinate
bool compareByX(const Point &a, const Point &b) {
	return a.x < b.x || (a.x == b.x && a.y < b.y);
}

// Compare points by y-coordinate
bool compareByY(const Point &a, const Point &b) {
	return a.y < b.y;
}

// Calculate Euclidean distance between two points
inline double distance(const Point &a, const Point &b) {
	return std::sqrt((a.x - b.x) * (a.x - b.x) + (a.y - b.y) * (a.y - b.y));
}

// A helper function to find the closest pair of points in a strip
std::pair<Point, Point> closestInStrip(std::vector<Point> &strip, double d) {
	double minDist = d;
	std::pair<Point, Point> closestPair;

	// Sort strip points according to y coordinate
	std::sort(strip.begin(), strip.end(), compareByY);

	// Pick all points one by one and try the next points till the difference
	// between y coordinates is smaller than d.
	// This is a proven fact that this loop runs at most 6 times
	for (size_t i = 0; i < strip.size(); ++i) {
		for (size_t j = i + 1; j < strip.size() && (strip[j].y - strip[i].y) < minDist; ++j) {
			double dist = distance(strip[i], strip[j]);
			if (dist < minDist) {
				minDist = dist;
				closestPair = {strip[i], strip[j]};
			}
		}
	}

	return closestPair;
}

// Recursive function to find the closest pair of points
std::pair<Point, Point> closestPairRec(std::vector<Point> &points, int left, int right) {
	if (right - left <= 3) {
		double minDist = std::numeric_limits<double>::max();
		std::pair<Point, Point> closestPair;

		for (int i = left; i < right; ++i) {
			for (int j = i + 1; j < right; ++j) {
				double dist = distance(points[i], points[j]);
				if (dist < minDist) {
					minDist = dist;
					closestPair = {points[i], points[j]};
				}
			}
		}

		std::sort(points.begin() + left, points.begin() + right, compareByY);
		return closestPair;
	}

	int mid = left + (right - left) / 2;
	Point midPoint = points[mid];

	auto leftClosest = closestPairRec(points, left, mid);
	auto rightClosest = closestPairRec(points, mid, right);

	double leftDist = distance(leftClosest.first, leftClosest.second);
	double rightDist = distance(rightClosest.first, rightClosest.second);

	std::pair<Point, Point> closestPair = (leftDist < rightDist) ? leftClosest : rightClosest;
	double minDist = std::min(leftDist, rightDist);

	std::vector<Point> strip;
	for (int i = left; i < right; ++i) {
		if (std::abs(points[i].x - midPoint.x) < minDist) {
			strip.push_back(points[i]);
		}
	}

	auto stripClosest = closestInStrip(strip, minDist);
	if (!stripClosest.first.x && !stripClosest.first.y && !stripClosest.second.x && !stripClosest.second.y) {
		return closestPair;
	}

	double stripDist = distance(stripClosest.first, stripClosest.second);
	return (stripDist < minDist) ? stripClosest : closestPair;
}

void processTestCase(std::vector<Point> &points) {
	std::sort(points.begin(), points.end(), compareByX);
	auto closestPair = closestPairRec(points, 0, points.size());

	printf("%.2f %.2f %.2f %.2f\n",
		   closestPair.first.x, closestPair.first.y,
		   closestPair.second.x, closestPair.second.y);
}

int main() {
	int n;
	while (std::cin >> n && n != 0) {
		std::vector<Point> points(n);
		for (int i = 0; i < n; ++i) {
			std::cin >> points[i].x >> points[i].y;
			points[i].id = i;
		}
		processTestCase(points);
	}
	return 0;
}
