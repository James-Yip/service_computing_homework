package scenario

import (
	"fmt"
	"time"

	"github.com/James-Yip/service_computing_homework/async-httpclient/agency"
)

func Synchronous() [10]agency.Response {
	beginning := time.Now()
	customer := agency.GetCustomerDetails()
	destinations := agency.GetRecommendedDestinations(customer)
	var responses [10]agency.Response
	for i, d := range destinations {
		q := agency.GetQuote(d)
		w := agency.GetWeatherForcast(d)
		responses[i] = agency.Response{
			Weather:     w,
			Destination: d,
			Quote:       q}
	}
	fmt.Println(time.Since(beginning))
	return responses
}
