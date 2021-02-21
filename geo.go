// Copyright Frontware International
// This package is used by several Frontware projects to handle basic tasks about geo location

package geo

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

//go:generate ffjson geo.go

type (
	// Nominatim is address structure returned by nominatim API.
	Nominatim struct {
		DisplayName string   `json:"display_name"`
		Address     *Address `json:"address"`
	}

	// Address use to query Mapstreet for reverse geo location
	Address struct {
		Country  string `json:"country_code"`
		Road     string `json:"road"`
		City     string `json:"city"`
		Postcode string `json:"postcode"`
		Region   string `json:"state"`
	}

	// Place is the struct of a geo place from nominatim.
	Place struct {
		Lat         float64 `json:"lat"`
		Long        float64 `json:"lon"`
		PlaceID     string  `json:"place_id"`
		DisplayName string  `json:"display_name"`
		Class       string  `json:"class"`
		Type        string  `json:"type"`
		Importance  float64 `json:"importance"`
		OSMType     string  `json:"osm_type"`
	}
)

// Distance function returns the distance (in meters) between two points of
// a given longitude and latitude relatively accurately (using a spherical
// approximation of the Earth) through the Haversin Distance Formula for
// great arc distance on a sphere with accuracy for small distances
// point coordinates are supplied in degrees and converted into rad. in the func
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
//  Here we get the distance
//  geo.Distance(13.76,100.50, 13.89, 101.12)
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// Reverse returns location name based on coordinates from openstreetmap API
// We wait 1 second before start because there is a rate limitation of 1 request per second
//  Reverse(13.7665269,100.6068431)
func Reverse(lat, lon float64) (address Nominatim, err error) {
	// curl "https://nominatim.openstreetmap.org/reverse?format=json&lat=18.8094923&lon=98.968031&zoom=18&addressdetails=1"

	baseURL := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=18&addressdetails=1", lat, lon)

	// We wait 1 second because terms of usage limit to 1 call / second (https://operations.osmfoundation.org/policies/nominatim/)
	time.Sleep(1 * time.Second)
	// Set a 10 seconds timeout to avoid keeping too many open sockets
	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	res, err := client.Get(baseURL)
	if err != nil {
		return
	}
	defer func() {
		res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	err = address.UnmarshalJSON(body)
	return
}

// GeoLocate returns coordinates based on address
//   GeoLocate(geo.Address{City:"Bangkok","Road":"Latprao 94, Town in Town",PostCode:10310})
func GeoLocate(address Address) (lat, long float64) {
	// curl "https://nominatim.openstreetmap.org/search/query?city=ottignies&street=pinchart 31&format=json

	query := url.QueryEscape(fmt.Sprintf("format=json&city=%s&street=%s&postcode=%s",
		address.City,
		address.Road,
		address.Postcode,
	))

	baseURL := "https://nominatim.openstreetmap.org/search/query?" + query

	// We wait 1 second because terms of usage limit to 1 call / second (https://operations.osmfoundation.org/policies/nominatim/)
	time.Sleep(1 * time.Second)
	// Set a 10 seconds timeout to avoid keeping too many open sockets
	client := http.Client{Timeout: time.Duration(10 * time.Second)}
	res, err := client.Get(baseURL)

	if err != nil {
		return
	}

	defer func() {
		res.Body.Close()
	}()

	var places []Place
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}

	if err = ffjson.Unmarshal(body, &places); err == nil && len(places) > 0 {
		lat = places[0].Lat
		long = places[0].Long
	}

	return
}

// hsin haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
