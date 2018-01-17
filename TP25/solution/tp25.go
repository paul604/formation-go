package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"formation-go/TP25/solution/yahoo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var woeid = map[string]string{
	"Paris":  "615702",
	"Cannes": "582935",
	"Rome":   "721943",
}

func main() {
	// http.HandleFunc("/", handler)           // each request calls handler
	// http.HandleFunc("/weather/*", weatherAPI) // each request calls handler
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homeHandler)
	myRouter.HandleFunc("/cities/", citiesHandler)
	myRouter.HandleFunc("/cities/{name}", cityHandler)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func citiesHandler(w http.ResponseWriter, r *http.Request) {
	data := []yahoo.Weather{}
	for _, id := range woeid {
		j, err := getCityWeather(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		data = append(data, j)
	}

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("An error occurred while marshaling cities weather in json format: ", err)
		return
	}

	fmt.Fprintf(w, "%s", j)

}

func cityHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	city := vars["name"]

	id, ok := woeid[city]
	if !ok {
		fmt.Println("Unsupported city: ", city)
		return
	}

	data, err := getCityWeather(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("An error occurred while marshaling weather in json format: ", err)
		return
	}

	fmt.Fprintf(w, "%s", j)
}

func getCityWeather(woeid string) (yahoo.Weather, error) {
	w, err := yahoo.GetWeatherByWoeid(woeid)
	if err != nil {
		return yahoo.Weather{}, errors.New("An error occurred calling the weather API:" + err.Error())
	}

	return w, nil
}
