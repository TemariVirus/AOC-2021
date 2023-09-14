package main

import (
	"regexp"
	"strconv"
	"strings"
)

type RebootStep struct {
	On         bool
	Start, End Point3D
}

type RebootRegion struct {
	Start, End Point3D
}

func (r RebootRegion) volume() int64 {
	return int64(r.End.X-r.Start.X+1) * int64(r.End.Y-r.Start.Y+1) * int64(r.End.Z-r.Start.Z+1)
}

func parseInput22(input string) []RebootStep {
	steps := []RebootStep{}
	for _, line := range strings.Split(input, "\n") {
		exp := regexp.MustCompile(`^(on|off) x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
		match := exp.FindStringSubmatch(line)[1:]
		on := match[0] == "on"
		coords := apply(match[1:], func(s string) int { return unwrap(strconv.Atoi(s)) })
		steps = append(steps, RebootStep{
			On:    on,
			Start: Point3D{min(coords[0], coords[1]), min(coords[2], coords[3]), min(coords[4], coords[5])},
			End:   Point3D{max(coords[0], coords[1]), max(coords[2], coords[3]), max(coords[4], coords[5])},
		})
	}
	return steps
}

func addRegion(regions []RebootRegion, step RebootStep) []RebootRegion {
	overlaps := []RebootRegion{}
	for i := 0; i < len(regions); i++ {
		region := regions[i]
		if step.Start.X <= region.End.X && step.End.X >= region.Start.X &&
			step.Start.Y <= region.End.Y && step.End.Y >= region.Start.Y &&
			step.Start.Z <= region.End.Z && step.End.Z >= region.Start.Z {
			overlaps = append(overlaps, region)
			regions = append(regions[:i], regions[i+1:]...)
			i--
		}
	}

	for _, overlap := range overlaps {
		if overlap.Start.X < step.Start.X {
			regions = append(regions, RebootRegion{overlap.Start, Point3D{step.Start.X - 1, overlap.End.Y, overlap.End.Z}})
			overlap = RebootRegion{Point3D{step.Start.X, overlap.Start.Y, overlap.Start.Z}, overlap.End}
		}
		if step.End.X < overlap.End.X {
			regions = append(regions, RebootRegion{Point3D{step.End.X + 1, overlap.Start.Y, overlap.Start.Z}, overlap.End})
			overlap = RebootRegion{overlap.Start, Point3D{step.End.X, overlap.End.Y, overlap.End.Z}}
		}

		if overlap.Start.Y < step.Start.Y {
			regions = append(regions, RebootRegion{overlap.Start, Point3D{overlap.End.X, step.Start.Y - 1, overlap.End.Z}})
			overlap = RebootRegion{Point3D{overlap.Start.X, step.Start.Y, overlap.Start.Z}, overlap.End}
		}
		if step.End.Y < overlap.End.Y {
			regions = append(regions, RebootRegion{Point3D{overlap.Start.X, step.End.Y + 1, overlap.Start.Z}, overlap.End})
			overlap = RebootRegion{overlap.Start, Point3D{overlap.End.X, step.End.Y, overlap.End.Z}}
		}

		if overlap.Start.Z < step.Start.Z {
			regions = append(regions, RebootRegion{overlap.Start, Point3D{overlap.End.X, overlap.End.Y, step.Start.Z - 1}})
			overlap = RebootRegion{Point3D{overlap.Start.X, overlap.Start.Y, step.Start.Z}, overlap.End}
		}
		if step.End.Z < overlap.End.Z {
			regions = append(regions, RebootRegion{Point3D{overlap.Start.X, overlap.Start.Y, step.End.Z + 1}, overlap.End})
			overlap = RebootRegion{overlap.Start, Point3D{overlap.End.X, overlap.End.Y, step.End.Z}}
		}
	}

	region := RebootRegion{step.Start, step.End}
	if step.On && region.volume() > 0 {
		regions = append(regions, region)
	}

	return regions
}

func countOn(regions []RebootRegion) int64 {
	return aggregate(regions, int64(0), func(agg int64, value RebootRegion, _ int) int64 {
		return agg + value.volume()
	})
}

func solution22Part1(input string) int64 {
	steps := parseInput22(input)
	regions := []RebootRegion{}
	for _, step := range steps {
		if step.Start.X < -50 || step.Start.X > 50 {
			continue
		}
		regions = addRegion(regions, step)
	}

	return countOn(regions)
}

func solution22Part2(input string) int64 {
	steps := parseInput22(input)
	regions := []RebootRegion{}
	for _, step := range steps {
		regions = addRegion(regions, step)
	}

	return countOn(regions)
}
