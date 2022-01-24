package cars

// UnitPrice is the cost to produce each car individually.
const UnitPrice = 10_000

// TenUnitsPrice is the cost to produce ten cars together.
const TenUnitsPrice = 95_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var tens uint = uint(carsCount) / 10
	var units uint = uint(carsCount) % 10
	return (tens * TenUnitsPrice) + (units * UnitPrice)
}
