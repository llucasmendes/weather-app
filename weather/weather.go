package weather

import (
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func getCityFromZipcode(zipcode string) (string, error) {
	client := resty.New()
	resp, err := client.R().
		SetResult(&map[string]interface{}{}).
		Get("https://viacep.com.br/ws/" + zipcode + "/json/")
	if err != nil {
		return "", err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return "", errors.New("can not find zipcode")
	}

	data := *(resp.Result().(*map[string]interface{}))
	city, exists := data["localidade"].(string)
	if !exists {
		return "", errors.New("invalid response from viaCEP")
	}

	return city, nil
}

func getTemperature(city string) (float64, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("key", "afc097e924ac4122b6d153241242405").
		SetQueryParam("q", city).
		SetResult(&map[string]interface{}{}).
		Get("http://api.weatherapi.com/v1/current.json")
	if err != nil {
		return 0, err
	}

	data := *(resp.Result().(*map[string]interface{}))
	current, exists := data["current"].(map[string]interface{})
	if !exists {
		return 0, errors.New("invalid response from WeatherAPI")
	}

	tempC, exists := current["temp_c"].(float64)
	if !exists {
		return 0, errors.New("invalid temperature data")
	}

	return tempC, nil
}
