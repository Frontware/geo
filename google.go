package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

//go:generate ffjson google.go

type (
	gpReview struct {
		ID      string `json:"_id" bson:"_id"` // Hash of author, venueid + time
		Aspects []struct {
			Rating int64  `json:"rating" bson:"rating"`
			Type   string `json:"type" bson:"type"`
		} `json:"aspects" bson:"aspects"`
		AuthorName      string     `json:"author_name" bson:"author_name"`
		AuthorURL       string     `json:"author_url" bson:"author_url"`
		Language        string     `json:"language" bson:"lg"`
		ProfilePhotoURL string     `json:"profile_photo_url" bson:"profile_photo_url"`
		Rating          int64      `json:"rating" bson:"rating"`
		Text            string     `json:"text" bson:"text,omitempty"`
		Time            *time.Time `json:"time,omitempty" bson:"time,omitempty"`
		VenueID         string     `json:"venueid" bson:"venueid"`
		ToShow          bool       `json:"toShow,omitempty" bson:"toShow,omitempty"` // Set to true if we have to show it on Venue page
	}

	// GoogleOpeningHours
	GoogleOpeningHours struct {
		OpenNow bool `json:"open_now"` // is a boolean value indicating if the place is open at the current time
		// WeekdayText is an array of seven strings representing the formatted opening hours for each day of the week.
		// If a language parameter was specified in the Place Details request,
		// the Places Service will format and localize the opening hours appropriately for that language.
		// The ordering of the elements in this array depends on the language parameter.
		// Some languages start the week on Monday while others start on Sunday.
		WeekdayText []string `json:"weekday_text"`
		//Periods is an array of opening periods covering seven days, starting from Sunday, in chronological order.
		Periods []struct {
			Close struct {
				Day  int    `json:"day"`  // day a number from 0–6, corresponding to the days of the week, starting on Sunday. For example, 2 means Tuesday.
				Time string `json:"time"` // time may contain a time of day in 24-hour hhmm format. Values are in the range 0000–2359. The time will be reported in the place’s time zone.
			} `json:"close"`
			Open struct {
				Day  int    `json:"day"`
				Time string `json:"time"`
			} `json:"open"`
		}
	}

	GooglePlace struct {
		Geometry struct {
			LocationType string `json:"location_type "` // APPROXIMATE, RANGE_INTERPOLATED
			Location     struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
		Icon         string             `json:"icon"`
		ID           string             `json:"id"`
		Name         string             `json:"name"`
		OpeningHours GoogleOpeningHours `json:"opening_hours"`
		Photos       []struct {
			Height           int64    `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int64    `json:"width"`
		} `json:"photos"`
		//PlaceID A textual identifier that uniquely identifies a place.
		// To retrieve information about the place, pass this identifier in the placeId field of a Places API request.
		// For more information about place IDs: https://developers.google.com/places/web-service/place-id
		PlaceID string `json:"place_id"`
		//price_level — The price level of the place, on a scale of 0 to 4. The exact amount indicated by a specific value will vary from region to region. Price levels are interpreted as follows:
		//0 — Free
		//1 — Inexpensive
		//2 — Moderate
		//3 — Expensive
		//4 — Very Expensive
		PriceLevel        int         `json:"price_level"`
		Rating            interface{} `json:"rating"`
		Reference         string      `json:"reference"`
		Scope             string      `json:"scope"`
		Types             []string    `json:"types"`
		PartialMatch      bool        `json:"partial_match"`
		Vicinity          string      `json:"vicinity"`
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		AdrAddress               string `json:"adr_address"`
		FormattedAddress         string `json:"formatted_address"`
		FormattedPhoneNumber     string `json:"formatted_phone_number"`
		InternationalPhoneNumber string `json:"international_phone_number"`
		//PermanentlyClosed is a boolean flag indicating whether the place has permanently shut down (value true).
		// If the place is not permanently closed, the flag is absent from the response.
		PermanentlyClosed bool       `json:"permanently_closed"`
		Reviews           []gpReview `json:"reviews"`
		URL               string     `json:"url"`
		UserRatingsTotal  int64      `json:"user_ratings_total"`
		UtcOffset         int64      `json:"utc_offset"`
		Website           string     `json:"website,omitempty"`
	}
)

var googleAPI string

// SetGoogleAPI set Google API key
func SetGoogleAPI(key string) {
	googleAPI = key
}

// GeoCode gets coordinates based on address from Google Service.
//  GeoCode("Avenue Louise 24, Bruxelles, Belgium","en")
func GeoCode(address, lg string) (g GooglePlace, err error) {
	if address == "" || googleAPI == "" {
		err = errors.New("Missing")
		return
	}

	if len(lg) != 2 {
		lg = "en"
	}

	// https://maps.google.com/maps/api/geocode/json?address=Ferme%20des%20Poursaude%2008420%20Villers-le-tilleul%20france

	address = url.QueryEscape(address)

	baseURL := "https://maps.google.com/maps/api/geocode/json?address="

	url := baseURL + address + "&hl=" + lg + "&key=" + googleAPI

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; <Android Version>; <Build Tag etc.>) AppleWebKit/<WebKit Rev> (KHTML, like Gecko) Chrome/<Chrome Rev> Mobile Safari/<WebKit Rev>")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return g, err
	}

	var body = &bytes.Buffer{}
	_, err = body.ReadFrom(res.Body)
	if err != nil {
		return g, err
	}
	res.Body.Close()

	var result struct {
		Results []GooglePlace `json:"results"`

		// OK indicates that no errors occurred; the place was successfully detected and at least one result was returned.
		// UNKNOWN_ERROR indicates a server-side error; trying again may be successful.
		// ZERO_RESULTS indicates that the reference was valid but no longer refers to a valid result. This may occur if the establishment is no longer in business.
		// OVER_QUERY_LIMIT indicates that you are over your quota.
		// REQUEST_DENIED indicates that your request was denied, generally because of lack of an invalid key parameter.
		// INVALID_REQUEST generally indicates that the query (reference) is missing.
		// NOT_FOUND indicates that the referenced location was not found in the Places database.
		Status   string `json:"status"`
		ErrorMsg string `json:"error_message"`
	}

	if err = json.Unmarshal(body.Bytes(), &result); err != nil {
		return g, err
	}
	if result.Status != "OK" {
		return g, errors.New(result.Status)
	}
	g = result.Results[0]
	return g, nil
}
