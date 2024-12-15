package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Particle struct {
	x int
	y int
}

func wrapIndex(value, maxValue int) int {
	return (value%maxValue + maxValue) % maxValue
}

func quadrant(x int, y int, width int, height int) int {
	centreX := (width - 1) / 2
	centreY := (height - 1) / 2

	if x < centreX {
		if y < centreY {
			return 1
		} else if y > centreY {
			return 2
		}
	} else if x > centreX {
		if y < centreY {
			return 3
		} else if y > centreY {
			return 4
		}
	}
	return 0
}

func safetyFactor(quadrantCounts map[int]int) int {
	value := 1
	for _, i := range []int{1, 2, 3, 4} {
		if v, ok := quadrantCounts[i]; ok {
			value *= v
		}
	}
	return value
}

func simulate(particles string, width int, height int, timesteps int) ([]Particle, map[int]int) {
	var newParticles []Particle
	quadrantCounts := make(map[int]int)

	regex, _ := regexp.Compile(`p=(-*\d+),(-*\d+) v=(-*\d+),(-*\d+)`)
	for _, particle := range strings.Split(particles, "\n") {
		// New particle locations
		line := regex.FindAllStringSubmatch(particle, -1)[0]
		x, _ := strconv.Atoi(line[1])
		y, _ := strconv.Atoi(line[2])
		vx, _ := strconv.Atoi(line[3])
		vy, _ := strconv.Atoi(line[4])
		xNew := wrapIndex(x+vx*timesteps, width)
		yNew := wrapIndex(y+vy*timesteps, height)
		newParticles = append(newParticles, Particle{xNew, yNew})

		// Quadrant
		quadrant := quadrant(xNew, yNew, width, height)
		if _, ok := quadrantCounts[quadrant]; ok {
			quadrantCounts[quadrant]++
		} else {
			quadrantCounts[quadrant] = 1
		}
	}

	return newParticles, quadrantCounts
}

func renderPNG(particles []Particle, width int, height int, filename string) error {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if slices.Contains(particles, Particle{x, y}) {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return err
	}

	return nil
}

func part1(particles string, width int, height int, timesteps int) int {
	_, quadrantCounts := simulate(particles, width, height, timesteps)
	return safetyFactor(quadrantCounts)
}

func part2(particles string, width int, height int) int {
	for timesteps := 0; timesteps < 10000; timesteps += 1 {
		newParticles, _ := simulate(strings.Clone(particles), width, height, timesteps)
		renderPNG(newParticles, width, height, fmt.Sprintf("tmp/%04d.png", timesteps))
	}
	return 0
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	width := 101
	height := 103
	timesteps := 100
	fmt.Printf("Part 1: %d\n", part1(string(file), width, height, timesteps))
	part2(strings.Clone(string(file)), width, height)
}
