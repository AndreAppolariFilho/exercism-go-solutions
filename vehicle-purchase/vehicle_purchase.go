package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var chosenVehicle string
	if option1 > option2 {
		chosenVehicle = option2
	} else {
		chosenVehicle = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", chosenVehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var tax float64 = 0.5
	if age <= 3 {
		tax = 0.80
	} else if age > 3 && age < 10 {
		tax = 0.70
	}
	return originalPrice * tax
}
