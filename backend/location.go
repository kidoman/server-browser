package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GeoLocation struct {
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
}

func location(ip string) (string, error) {
	url := fmt.Sprintf("https://freegeoip.net/json/%v", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var g GeoLocation
	if err := json.Unmarshal(body, &g); err != nil {
		return "", err
	}
	if g.City != "" {
		return g.City, nil
	}
	if g.CountryCode != "" {
		return g.CountryCode, nil
	}
	return "", fmt.Errorf("could not resolve location for %v", ip)
}
