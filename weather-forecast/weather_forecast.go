// Package weather is simply dummy text of the printing and typesetting industry.
package weather

// CurrentCondition is simply dummy text of the printing and typesetting industry.
var CurrentCondition string

// CurrentLocation is simply dummy text of the printing and typesetting industry.
var CurrentLocation string

// Forecast is simply dummy text of the printing and typesetting industry.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
