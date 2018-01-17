package yahoo

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleGetWoeidByCityName() {
	woeid, _ := GetWoeidByCityName("Rome, IT")
	fmt.Printf("%s\n", woeid)

	woeid, _ = GetWoeidByCityName("Paris, FR")
	fmt.Printf("%s\n", woeid)

	woeid, _ = GetWoeidByCityName("Cannes, PACA, FR")
	fmt.Printf("%s\n", woeid)

	// Output:
	// 721943
	// 615702
	// 582935
}

func TestGetWheaterByWoeid(t *testing.T) {

	romeWoeid := "721943"
	w, err := GetWeatherByWoeid(romeWoeid)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	cityID := w.Woeid
	desc := w.Description
	temp := w.Temperature
	
	if cityID != romeWoeid {
		fmt.Println("Wrong city woeid:" + cityID)
		t.Fail()
	}

	if temp > 50 || temp < -15 {
		fmt.Println("Weird temperature value: " + strconv.Itoa(temp) + " C")
		t.Fail()
	}

	if len(desc) < 3 {
		fmt.Println("Weird weather description: " + desc)
		t.Fail()
	}
}
