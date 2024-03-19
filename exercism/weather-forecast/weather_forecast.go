// Package weather provides information 
// about the weather in a determined location.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the current location.
var CurrentLocation string

// Forecast returns a weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
