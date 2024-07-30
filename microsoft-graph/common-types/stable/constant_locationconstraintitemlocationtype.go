package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraintItemLocationType string

const (
	LocationConstraintItemLocationType_BusinessAddress LocationConstraintItemLocationType = "businessAddress"
	LocationConstraintItemLocationType_ConferenceRoom  LocationConstraintItemLocationType = "conferenceRoom"
	LocationConstraintItemLocationType_Default         LocationConstraintItemLocationType = "default"
	LocationConstraintItemLocationType_GeoCoordinates  LocationConstraintItemLocationType = "geoCoordinates"
	LocationConstraintItemLocationType_HomeAddress     LocationConstraintItemLocationType = "homeAddress"
	LocationConstraintItemLocationType_Hotel           LocationConstraintItemLocationType = "hotel"
	LocationConstraintItemLocationType_LocalBusiness   LocationConstraintItemLocationType = "localBusiness"
	LocationConstraintItemLocationType_PostalAddress   LocationConstraintItemLocationType = "postalAddress"
	LocationConstraintItemLocationType_Restaurant      LocationConstraintItemLocationType = "restaurant"
	LocationConstraintItemLocationType_StreetAddress   LocationConstraintItemLocationType = "streetAddress"
)

func PossibleValuesForLocationConstraintItemLocationType() []string {
	return []string{
		string(LocationConstraintItemLocationType_BusinessAddress),
		string(LocationConstraintItemLocationType_ConferenceRoom),
		string(LocationConstraintItemLocationType_Default),
		string(LocationConstraintItemLocationType_GeoCoordinates),
		string(LocationConstraintItemLocationType_HomeAddress),
		string(LocationConstraintItemLocationType_Hotel),
		string(LocationConstraintItemLocationType_LocalBusiness),
		string(LocationConstraintItemLocationType_PostalAddress),
		string(LocationConstraintItemLocationType_Restaurant),
		string(LocationConstraintItemLocationType_StreetAddress),
	}
}

func (s *LocationConstraintItemLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocationConstraintItemLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocationConstraintItemLocationType(input string) (*LocationConstraintItemLocationType, error) {
	vals := map[string]LocationConstraintItemLocationType{
		"businessaddress": LocationConstraintItemLocationType_BusinessAddress,
		"conferenceroom":  LocationConstraintItemLocationType_ConferenceRoom,
		"default":         LocationConstraintItemLocationType_Default,
		"geocoordinates":  LocationConstraintItemLocationType_GeoCoordinates,
		"homeaddress":     LocationConstraintItemLocationType_HomeAddress,
		"hotel":           LocationConstraintItemLocationType_Hotel,
		"localbusiness":   LocationConstraintItemLocationType_LocalBusiness,
		"postaladdress":   LocationConstraintItemLocationType_PostalAddress,
		"restaurant":      LocationConstraintItemLocationType_Restaurant,
		"streetaddress":   LocationConstraintItemLocationType_StreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationConstraintItemLocationType(input)
	return &out, nil
}
