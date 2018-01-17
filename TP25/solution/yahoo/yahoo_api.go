package yahoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Weather type
type Weather struct {
	Woeid       string `json:"CityWoeid"`
	Description string `json:"WeatherDescription"`
	Temperature int    `json:"DegreeCelsiusTemperature"`
}

type woeidQueryResult struct {
	Query struct {
		Lang    string
		Count   int
		Created string
		Results struct {
			Place struct {
				Woeid string
			}
		}
	}
}

type yahooQueryResult struct {
	Query struct {
		Lang    string
		Count   int
		Created string
		Results struct {
			Channel struct {
				Item struct {
					Condition struct {
						Temp string
						Text string
						Code string
						Date string
					}
				}
			}
		}
	}
}

// GetWoeidByCityName retrieves Yahoo woeid
// (Where On Earth Identifier) for a given cityName.
func GetWoeidByCityName(cityName string) (string, error) {

	url := buildWoeidQuery(cityName)

	res, err := makeHTTPRequest(url)
	if err != nil {
		return "", err
	}

	yr, err := unmarshalWoeidResult(res)
	if err != nil {
		return "", err
	}

	return yr.Query.Results.Place.Woeid, nil
}

// GetWeatherByWoeid get current Weather from Yahoo
// for a given woeid.
func GetWeatherByWoeid(cityID string) (Weather, error) {

	url := buildWeatherQuery(cityID)

	res, err := makeHTTPRequest(url)
	if err != nil {
		return Weather{}, err
	}

	yr, err := unmarshalWeatherResult(res)
	if err != nil {
		return Weather{}, err
	}

	description := yr.Query.Results.Channel.Item.Condition.Text
	temperature, err := strconv.Atoi(yr.Query.Results.Channel.Item.Condition.Temp)

	return Weather{
		cityID,
		description,
		temperature,
	}, nil
}


func buildWoeidQuery(cityName string) string {
	query := "select woeid from geo.places where text='" + cityName + "'"
	return buildQuery(query)
}

func buildWeatherQuery(cityID string) string {
	query := "select item.condition from weather.forecast where woeid="+cityID+" and u='c'"
	return buildQuery(query)
}

func buildQuery(query string) string {

	u := url.URL{
		Scheme: "http",
		Host:   "query.yahooapis.com",
		Path:   "v1/public/yql",
	}

	q := url.Values{}
	q.Set("q", query)
	q.Set("format", "json")
	u.RawQuery = q.Encode()
	return u.String()
}

func makeHTTPRequest(url string) ([]byte, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func unmarshalWeatherResult(data []byte) (yahooQueryResult, error) {

	var result yahooQueryResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("Error unmashaling the following json string: %s\n", data)
		return yahooQueryResult{}, err
	}
	return result, nil
}

func unmarshalWoeidResult(data []byte) (woeidQueryResult, error) {

	var result woeidQueryResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("Error unmashaling the following json string: %s\n", data)
		return woeidQueryResult{}, err
	}
	return result, nil
}

