// Package weather utility.
package weather

// CurrentCondition: condition for forecasting.
var CurrentCondition string

// CurrentLocation: location.
var CurrentLocation string

// Forecast the weather of the location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
