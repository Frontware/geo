// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: geo.go

package geo

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Address) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Address) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"country_code":`)
	fflib.WriteJsonString(buf, string(j.Country))
	buf.WriteString(`,"road":`)
	fflib.WriteJsonString(buf, string(j.Road))
	buf.WriteString(`,"city":`)
	fflib.WriteJsonString(buf, string(j.City))
	buf.WriteString(`,"postcode":`)
	fflib.WriteJsonString(buf, string(j.Postcode))
	buf.WriteString(`,"state":`)
	fflib.WriteJsonString(buf, string(j.Region))
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAddressbase = iota
	ffjtAddressnosuchkey

	ffjtAddressCountry

	ffjtAddressRoad

	ffjtAddressCity

	ffjtAddressPostcode

	ffjtAddressRegion
)

var ffjKeyAddressCountry = []byte("country_code")

var ffjKeyAddressRoad = []byte("road")

var ffjKeyAddressCity = []byte("city")

var ffjKeyAddressPostcode = []byte("postcode")

var ffjKeyAddressRegion = []byte("state")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Address) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Address) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAddressbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAddressnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyAddressCountry, kn) {
						currentKey = ffjtAddressCountry
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAddressCity, kn) {
						currentKey = ffjtAddressCity
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyAddressPostcode, kn) {
						currentKey = ffjtAddressPostcode
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyAddressRoad, kn) {
						currentKey = ffjtAddressRoad
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyAddressRegion, kn) {
						currentKey = ffjtAddressRegion
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyAddressRegion, kn) {
					currentKey = ffjtAddressRegion
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAddressPostcode, kn) {
					currentKey = ffjtAddressPostcode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAddressCity, kn) {
					currentKey = ffjtAddressCity
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAddressRoad, kn) {
					currentKey = ffjtAddressRoad
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyAddressCountry, kn) {
					currentKey = ffjtAddressCountry
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAddressnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAddressCountry:
					goto handle_Country

				case ffjtAddressRoad:
					goto handle_Road

				case ffjtAddressCity:
					goto handle_City

				case ffjtAddressPostcode:
					goto handle_Postcode

				case ffjtAddressRegion:
					goto handle_Region

				case ffjtAddressnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Country:

	/* handler: j.Country type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Country = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Road:

	/* handler: j.Road type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Road = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_City:

	/* handler: j.City type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.City = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Postcode:

	/* handler: j.Postcode type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Postcode = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Region:

	/* handler: j.Region type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Region = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *Nominatim) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Nominatim) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"display_name":`)
	fflib.WriteJsonString(buf, string(j.DisplayName))
	if j.Address != nil {
		buf.WriteString(`,"address":`)

		{

			err = j.Address.MarshalJSONBuf(buf)
			if err != nil {
				return err
			}

		}
	} else {
		buf.WriteString(`,"address":null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtNominatimbase = iota
	ffjtNominatimnosuchkey

	ffjtNominatimDisplayName

	ffjtNominatimAddress
)

var ffjKeyNominatimDisplayName = []byte("display_name")

var ffjKeyNominatimAddress = []byte("address")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Nominatim) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Nominatim) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtNominatimbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtNominatimnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyNominatimAddress, kn) {
						currentKey = ffjtNominatimAddress
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyNominatimDisplayName, kn) {
						currentKey = ffjtNominatimDisplayName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyNominatimAddress, kn) {
					currentKey = ffjtNominatimAddress
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyNominatimDisplayName, kn) {
					currentKey = ffjtNominatimDisplayName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtNominatimnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtNominatimDisplayName:
					goto handle_DisplayName

				case ffjtNominatimAddress:
					goto handle_Address

				case ffjtNominatimnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_DisplayName:

	/* handler: j.DisplayName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.DisplayName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Address:

	/* handler: j.Address type=geo.Address kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Address = nil

		} else {

			if j.Address == nil {
				j.Address = new(Address)
			}

			err = j.Address.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// MarshalJSON marshal bytes to json - template
func (j *Place) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *Place) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"lat":`)
	fflib.AppendFloat(buf, float64(j.Lat), 'g', -1, 64)
	buf.WriteString(`,"lon":`)
	fflib.AppendFloat(buf, float64(j.Long), 'g', -1, 64)
	buf.WriteString(`,"place_id":`)
	fflib.WriteJsonString(buf, string(j.PlaceID))
	buf.WriteString(`,"display_name":`)
	fflib.WriteJsonString(buf, string(j.DisplayName))
	buf.WriteString(`,"class":`)
	fflib.WriteJsonString(buf, string(j.Class))
	buf.WriteString(`,"type":`)
	fflib.WriteJsonString(buf, string(j.Type))
	buf.WriteString(`,"importance":`)
	fflib.AppendFloat(buf, float64(j.Importance), 'g', -1, 64)
	buf.WriteString(`,"osm_type":`)
	fflib.WriteJsonString(buf, string(j.OSMType))
	buf.WriteByte('}')
	return nil
}

const (
	ffjtPlacebase = iota
	ffjtPlacenosuchkey

	ffjtPlaceLat

	ffjtPlaceLong

	ffjtPlacePlaceID

	ffjtPlaceDisplayName

	ffjtPlaceClass

	ffjtPlaceType

	ffjtPlaceImportance

	ffjtPlaceOSMType
)

var ffjKeyPlaceLat = []byte("lat")

var ffjKeyPlaceLong = []byte("lon")

var ffjKeyPlacePlaceID = []byte("place_id")

var ffjKeyPlaceDisplayName = []byte("display_name")

var ffjKeyPlaceClass = []byte("class")

var ffjKeyPlaceType = []byte("type")

var ffjKeyPlaceImportance = []byte("importance")

var ffjKeyPlaceOSMType = []byte("osm_type")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Place) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Place) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtPlacebase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtPlacenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffjKeyPlaceClass, kn) {
						currentKey = ffjtPlaceClass
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyPlaceDisplayName, kn) {
						currentKey = ffjtPlaceDisplayName
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyPlaceImportance, kn) {
						currentKey = ffjtPlaceImportance
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyPlaceLat, kn) {
						currentKey = ffjtPlaceLat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyPlaceLong, kn) {
						currentKey = ffjtPlaceLong
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyPlaceOSMType, kn) {
						currentKey = ffjtPlaceOSMType
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyPlacePlaceID, kn) {
						currentKey = ffjtPlacePlaceID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyPlaceType, kn) {
						currentKey = ffjtPlaceType
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyPlaceOSMType, kn) {
					currentKey = ffjtPlaceOSMType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyPlaceImportance, kn) {
					currentKey = ffjtPlaceImportance
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyPlaceType, kn) {
					currentKey = ffjtPlaceType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyPlaceClass, kn) {
					currentKey = ffjtPlaceClass
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyPlaceDisplayName, kn) {
					currentKey = ffjtPlaceDisplayName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyPlacePlaceID, kn) {
					currentKey = ffjtPlacePlaceID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyPlaceLong, kn) {
					currentKey = ffjtPlaceLong
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyPlaceLat, kn) {
					currentKey = ffjtPlaceLat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtPlacenosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtPlaceLat:
					goto handle_Lat

				case ffjtPlaceLong:
					goto handle_Long

				case ffjtPlacePlaceID:
					goto handle_PlaceID

				case ffjtPlaceDisplayName:
					goto handle_DisplayName

				case ffjtPlaceClass:
					goto handle_Class

				case ffjtPlaceType:
					goto handle_Type

				case ffjtPlaceImportance:
					goto handle_Importance

				case ffjtPlaceOSMType:
					goto handle_OSMType

				case ffjtPlacenosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Lat:

	/* handler: j.Lat type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Lat = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Long:

	/* handler: j.Long type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Long = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PlaceID:

	/* handler: j.PlaceID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.PlaceID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DisplayName:

	/* handler: j.DisplayName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.DisplayName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Class:

	/* handler: j.Class type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Class = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: j.Type type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Type = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Importance:

	/* handler: j.Importance type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Importance = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OSMType:

	/* handler: j.OSMType type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.OSMType = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
