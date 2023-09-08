package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationLocationType string

const (
	LocationLocationTypebusinessAddress LocationLocationType = "BusinessAddress"
	LocationLocationTypeconferenceRoom  LocationLocationType = "ConferenceRoom"
	LocationLocationTypedefault         LocationLocationType = "Default"
	LocationLocationTypegeoCoordinates  LocationLocationType = "GeoCoordinates"
	LocationLocationTypehomeAddress     LocationLocationType = "HomeAddress"
	LocationLocationTypehotel           LocationLocationType = "Hotel"
	LocationLocationTypelocalBusiness   LocationLocationType = "LocalBusiness"
	LocationLocationTypepostalAddress   LocationLocationType = "PostalAddress"
	LocationLocationTyperestaurant      LocationLocationType = "Restaurant"
	LocationLocationTypestreetAddress   LocationLocationType = "StreetAddress"
)

func PossibleValuesForLocationLocationType() []string {
	return []string{
		string(LocationLocationTypebusinessAddress),
		string(LocationLocationTypeconferenceRoom),
		string(LocationLocationTypedefault),
		string(LocationLocationTypegeoCoordinates),
		string(LocationLocationTypehomeAddress),
		string(LocationLocationTypehotel),
		string(LocationLocationTypelocalBusiness),
		string(LocationLocationTypepostalAddress),
		string(LocationLocationTyperestaurant),
		string(LocationLocationTypestreetAddress),
	}
}

func parseLocationLocationType(input string) (*LocationLocationType, error) {
	vals := map[string]LocationLocationType{
		"businessaddress": LocationLocationTypebusinessAddress,
		"conferenceroom":  LocationLocationTypeconferenceRoom,
		"default":         LocationLocationTypedefault,
		"geocoordinates":  LocationLocationTypegeoCoordinates,
		"homeaddress":     LocationLocationTypehomeAddress,
		"hotel":           LocationLocationTypehotel,
		"localbusiness":   LocationLocationTypelocalBusiness,
		"postaladdress":   LocationLocationTypepostalAddress,
		"restaurant":      LocationLocationTyperestaurant,
		"streetaddress":   LocationLocationTypestreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationLocationType(input)
	return &out, nil
}
