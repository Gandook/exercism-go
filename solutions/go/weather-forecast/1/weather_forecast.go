// Package weather provides tools to report the weather condition.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the current location for the weather report.
var CurrentLocation string

// Forecast returns a string containing the weather report for the given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
