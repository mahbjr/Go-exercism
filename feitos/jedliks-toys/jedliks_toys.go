package jedlik

import "fmt"

// Note: Car is already defined in car.go, so we're just implementing methods for it

// Drive drives the car once. If there is not enough battery to drive one more time,
// the car will not move.
func (car *Car) Drive() {
    // Check if there's enough battery to drive
    if car.battery >= car.batteryDrain {
        car.distance += car.speed
        car.battery -= car.batteryDrain
    }
}

// DisplayDistance returns the distance as a string with a "Driven X meters" format.
func (car Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns the battery percentage as a string with a "Battery at X%" format.
func (car Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track distance.
func (car Car) CanFinish(trackDistance int) bool {
    // Calculate how many drives the car can make with the current battery
    maxDrives := car.battery / car.batteryDrain
    
    // Calculate how far the car can go with those drives
    maxDistance := maxDrives * car.speed
    
    // Check if the car can reach or exceed the track distance
    return maxDistance >= trackDistance
}
