package weather

import (
	"testing"
)

func TestGetCityFromZipcode(t *testing.T) {
	city, err := getCityFromZipcode("01001000")
	if err != nil {
		t.Errorf("Error getting city from zipcode: %v", err)
	}
	if city != "São Paulo" {
		t.Errorf("Expected São Paulo, got %s", city)
	}
}

func TestGetTemperature(t *testing.T) {
	temp, err := getTemperature("São Paulo")
	if err != nil {
		t.Errorf("Error getting temperature: %v", err)
	}
	if temp == 0 {
		t.Errorf("Expected non-zero temperature, got %f", temp)
	}
}
