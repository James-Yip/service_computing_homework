package agency

import "time"

type Weather struct{}

type Destination struct{}

type Quoting struct{}

type Customer struct{}

type Response struct {
	Weather     Weather
	Destination Destination
	Quote       Quoting
}

// provides information about an authenticated user of the travel agency.
func GetCustomerDetails() Customer {
	time.Sleep(150 * time.Millisecond)
	return Customer{}
}

// provides ten recommended destinations for an authenticated customer.
func GetRecommendedDestinations(c Customer) [10]Destination {
	time.Sleep(250 * time.Millisecond)
	return [10]Destination{}
}

// provides price calculation for a customer to travel to a recommended destination.
func GetQuote(d Destination) Quoting {
	time.Sleep(170 * time.Millisecond)
	return Quoting{}
}

// provides weather forecast for a given destination.
func GetWeatherForcast(d Destination) Weather {
	time.Sleep(330 * time.Millisecond)
	return Weather{}
}
