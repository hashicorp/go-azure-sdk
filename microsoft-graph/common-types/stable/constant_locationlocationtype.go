package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationLocationType string

const (
	LocationLocationType_BusinessAddress LocationLocationType = "businessAddress"
	LocationLocationType_ConferenceRoom  LocationLocationType = "conferenceRoom"
	LocationLocationType_Default         LocationLocationType = "default"
	LocationLocationType_GeoCoordinates  LocationLocationType = "geoCoordinates"
	LocationLocationType_HomeAddress     LocationLocationType = "homeAddress"
	LocationLocationType_Hotel           LocationLocationType = "hotel"
	LocationLocationType_LocalBusiness   LocationLocationType = "localBusiness"
	LocationLocationType_PostalAddress   LocationLocationType = "postalAddress"
	LocationLocationType_Restaurant      LocationLocationType = "restaurant"
	LocationLocationType_StreetAddress   LocationLocationType = "streetAddress"
)

func PossibleValuesForLocationLocationType() []string {
	return []string{
		string(LocationLocationType_BusinessAddress),
		string(LocationLocationType_ConferenceRoom),
		string(LocationLocationType_Default),
		string(LocationLocationType_GeoCoordinates),
		string(LocationLocationType_HomeAddress),
		string(LocationLocationType_Hotel),
		string(LocationLocationType_LocalBusiness),
		string(LocationLocationType_PostalAddress),
		string(LocationLocationType_Restaurant),
		string(LocationLocationType_StreetAddress),
	}
}

func (s *LocationLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationLocationType(input string) (*LocationLocationType, error) {
	vals := map[string]LocationLocationType{
		"businessaddress": LocationLocationType_BusinessAddress,
		"conferenceroom":  LocationLocationType_ConferenceRoom,
		"default":         LocationLocationType_Default,
		"geocoordinates":  LocationLocationType_GeoCoordinates,
		"homeaddress":     LocationLocationType_HomeAddress,
		"hotel":           LocationLocationType_Hotel,
		"localbusiness":   LocationLocationType_LocalBusiness,
		"postaladdress":   LocationLocationType_PostalAddress,
		"restaurant":      LocationLocationType_Restaurant,
		"streetaddress":   LocationLocationType_StreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationLocationType(input)
	return &out, nil
}
