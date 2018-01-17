package main

import (
	"fmt"
	"sort"
)

var worldCountries = []string{
	"Germany",
	"France",
	"United Kingdom",
	"Italy",
	"Spain",
	"Poland",
	"Romania",
	"Netherlands",
	"Belgium",
	"Nigeria",
	"Ethiopia",
	"Egypt",
	"Democratic Republic of the Congo",
	"South Africa",
	"Tanzania",
	"Kenya",
	"Algeria",
	"Sudan",
	"Uganda",
	"Morocco",
}

var worldCapitals = map[string]string{
	"Poland":      "Varsaw",
	"Romania":     "Bucharest",
	"Netherlands": "Amsterdam",
	"Belgium":     "Brussels",
	"Nigeria":     "Abuja",
	"Kenya":       "Nairobi",
	"Algeria":     "Algiers",
}

func main() {
	knownCountries := []string{}

	// Complete the following code
	for _, country := range worldCountries {
		// Iterate over all coutries in worldCountries
		if city, ok := worldCapitals[country]; ok {
			// Check if worldCapitals has a key that matches country :
			// 	use the declaration of an `ok` variable in the if condition
			fmt.Printf("We know %s's capital: it's %s 😎\n", country, city)
			knownCountries = append(knownCountries, country)
		} else {
			fmt.Printf("We don't know %s capital 😔\n", country)
		}
	}

	fmt.Println("\nWe know capitals for the following countries:")


	// Sort knownCountries : Use package sort of the standard library
	sort.Strings(knownCountries)

	for _, c := range knownCountries {
		fmt.Printf("\t%s 😎\n", c)
	}
}
