// Package weather is responsible for logic related with weather forecast.
package weather

// CurrentCondition: Condition about the weather.
var CurrentCondition string

// CurrentLocation: Location that you are doing your weather analysis.
var CurrentLocation string

// Forecast: This function return a string representation telling the current weather condition of a location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
