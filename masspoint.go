package main

// MassPoint is a body with position and mass information
type MassPoint struct {
	x, y, z, mass float64
}

func addMassPoints(a MassPoint, b MassPoint) MassPoint {
	return MassPoint{
		a.x + b.x,
		a.y + b.y,
		a.z + b.z,
		a.mass + b.mass,
	}
}

func toWeightedHyperspace(a MassPoint) MassPoint {
	return MassPoint{
		a.x * a.mass,
		a.y * a.mass,
		a.z * a.mass,
		a.mass,
	}
}
