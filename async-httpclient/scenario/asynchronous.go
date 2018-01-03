package scenario

import (
	"fmt"
	"time"

	"github.com/James-Yip/service_computing_homework/async-httpclient/agency"
)

func Asynchronous() [10]agency.Response {
	beginning := time.Now()
	customer := agency.GetCustomerDetails()
	destinations := agency.GetRecommendedDestinations(customer)
	var responses [10]agency.Response
	quotes := [10]chan agency.Quoting{}
	weathers := [10]chan agency.Weather{}

	for i := 0; i < 10; i++ {
		quotes[i] = make(chan agency.Quoting)
		weathers[i] = make(chan agency.Weather)
	}

	for i, d := range destinations {
		idx := i
		dest := d
		// responses[i].Destination = d
		go func() {
			// responses[i].Quote = agency.GetQuote(d)
			quotes[idx] <- agency.GetQuote(dest)
		}()

		go func() {
			// responses[i].Weather = agency.GetWeatherForcast(d)
			weathers[idx] <- agency.GetWeatherForcast(dest)
		}()
	}

	// combine the retrieved data
	for i, d := range destinations {
		responses[i] = agency.Response{
			Weather:     <-weathers[i],
			Destination: d,
			Quote:       <-quotes[i]}
	}

	fmt.Println(time.Since(beginning))
	return responses
}
