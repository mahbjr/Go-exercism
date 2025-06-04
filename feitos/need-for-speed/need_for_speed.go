package speed

// Car represents a remote controlled car with specific properties.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100, // Initial battery is always 100%
		distance:     0,   // Initial distance is always 0
	}
}

// Track represents a race track with a specific distance.
type Track struct {
	distance int
}

// NewTrack creates a new track with the specified distance.
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// Check if the car has enough battery to move
	if car.battery < car.batteryDrain {
		return car
	}

	// Update the car's properties
	car.battery -= car.batteryDrain
	car.distance += car.speed

	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// Calculate the remaining distance to finish the track
	remainingDistance := track.distance - car.distance

	// Calculate how far the car can go with the current battery
	maxDistance := (car.battery / car.batteryDrain) * car.speed

	// Check if the car can finish the track
	return maxDistance >= remainingDistance
}
