package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return choice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var factor float64
	if age < 3 {
		factor = 0.8
	} else if age < 10 {
		factor = 0.7
	} else {
		factor = 0.5
	}
	return originalPrice * factor
}
