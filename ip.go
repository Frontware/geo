package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//go:generate ffjson ip.go

type IPLocation struct {
	City          interface{} `json:"city"`
	ContinentCode string      `json:"continent_code"`
	ContinentName string      `json:"continent_name"`
	CountryCode   string      `json:"country_code"`
	CountryName   string      `json:"country_name"`
	IP            string      `json:"ip"`
	Latitude      int64       `json:"latitude"`
	Location      struct {
		CallingCode             string      `json:"calling_code"`
		Capital                 string      `json:"capital"`
		CountryFlag             string      `json:"country_flag"`
		CountryFlagEmoji        string      `json:"country_flag_emoji"`
		CountryFlagEmojiUnicode string      `json:"country_flag_emoji_unicode"`
		GeonameID               interface{} `json:"geoname_id"`
		IsEu                    bool        `json:"is_eu"`
		Languages               []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Native string `json:"native"`
		} `json:"languages"`
	} `json:"location"`
	Longitude  int64       `json:"longitude"`
	RegionCode interface{} `json:"region_code"`
	RegionName interface{} `json:"region_name"`
	Type       string      `json:"type"`
	Zip        interface{} `json:"zip"`
}

func LocateIP(ip string) (loc IPLocation, err error) {
	if ip == "" || ipstackeapi == "" {
		err = errors.New("Missing")
		return
	}
	// https://ipstack.com/documentation
	var baseURL = fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip, ipstackeapi)

	// Set a 5 seconds timeout to avoid keeping too many open sockets
	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	res, err := client.Get(baseURL)
	if err != nil {
		return
	}

	defer func() {
		res.Body.Close()
	}()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	return
}

// IPGeocode returns geo info based on IP
// Details https://rapidapi.com/apility.io/api/ip-geolocation
func IPGeocode(ip string) (err error) {
	if ip == "" || rapidapi == "" {
		err = errors.New("Missing")
		return
	}

	url := "https://apility-io-ip-geolocation-v1.p.rapidapi.com/" + ip
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "apility-io-ip-geolocation-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", rapidapi)
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// TODO : still have to return value
	fmt.Println(res)
	fmt.Println(string(body))
	return
}

type IPAPI struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryName        string  `json:"country_name"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  float64 `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
}

// GetLocationFromIP returns Location information based on IP
// More info https://ipapi.co/api/?go#introduction
func GetLocationFromIP(ip string) (ipapi IPAPI, err error) {
	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json/", ip), nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := ipapiClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = ipapi.UnmarshalJSON(body)
	return
}
