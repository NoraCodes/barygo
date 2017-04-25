package main

import (
	"fmt"
	"io"
	"os"
)

// Handle possible errors by aborting
func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// Close a file opened elsewhere
func closeFile(fi *os.File) {
	err := fi.Close()
	handle(err)
}

func main() {
	// TODO: Replace with actual selectable
	// Open the input file for reading
	file, err := os.Open("/tmp/bodies.txt")
	handle(err)
	defer closeFile(file)

	// A buffer for the MassPoints
	var masspoints []MassPoint

	// Scan the file for MassPoints
	var count int
	for {
		var x, y, z, mass float64
		_, err = fmt.Fscanf(file, "%f:%f:%f:%f", &x, &y, &z, &mass)
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}
		masspoints = append(masspoints, MassPoint{x, y, z, mass})

		count++
	}

	fmt.Printf("Loaded %d values from file.\n", count)

	// Map the points into the mass-weighed hyperspace
	hyperspace := make([]MassPoint, count)
	for _, v := range masspoints {
		hyperspace = append(hyperspace, toWeightedHyperspace(v))
	}

	// Add up all the points from hyperspace
	var systemHypercenter MassPoint
	for _, v := range hyperspace {
		systemHypercenter = addMassPoints(systemHypercenter, v)
	}

	// Pull the average out of the hyper space into real coordinates
	systemAverage := MassPoint{
		systemHypercenter.x / systemHypercenter.mass,
		systemHypercenter.y / systemHypercenter.mass,
		systemHypercenter.z / systemHypercenter.mass,
		systemHypercenter.mass,
	}

	fmt.Printf("System barycenter is at (%f, %f, %f) and the system's mass is %f.\n",
		systemAverage.x,
		systemAverage.y,
		systemAverage.z,
		systemAverage.mass)
}
