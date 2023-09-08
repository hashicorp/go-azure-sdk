package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraintItemLocationType string

const (
	LocationConstraintItemLocationTypebusinessAddress LocationConstraintItemLocationType = "BusinessAddress"
	LocationConstraintItemLocationTypeconferenceRoom  LocationConstraintItemLocationType = "ConferenceRoom"
	LocationConstraintItemLocationTypedefault         LocationConstraintItemLocationType = "Default"
	LocationConstraintItemLocationTypegeoCoordinates  LocationConstraintItemLocationType = "GeoCoordinates"
	LocationConstraintItemLocationTypehomeAddress     LocationConstraintItemLocationType = "HomeAddress"
	LocationConstraintItemLocationTypehotel           LocationConstraintItemLocationType = "Hotel"
	LocationConstraintItemLocationTypelocalBusiness   LocationConstraintItemLocationType = "LocalBusiness"
	LocationConstraintItemLocationTypepostalAddress   LocationConstraintItemLocationType = "PostalAddress"
	LocationConstraintItemLocationTyperestaurant      LocationConstraintItemLocationType = "Restaurant"
	LocationConstraintItemLocationTypestreetAddress   LocationConstraintItemLocationType = "StreetAddress"
)

func PossibleValuesForLocationConstraintItemLocationType() []string {
	return []string{
		string(LocationConstraintItemLocationTypebusinessAddress),
		string(LocationConstraintItemLocationTypeconferenceRoom),
		string(LocationConstraintItemLocationTypedefault),
		string(LocationConstraintItemLocationTypegeoCoordinates),
		string(LocationConstraintItemLocationTypehomeAddress),
		string(LocationConstraintItemLocationTypehotel),
		string(LocationConstraintItemLocationTypelocalBusiness),
		string(LocationConstraintItemLocationTypepostalAddress),
		string(LocationConstraintItemLocationTyperestaurant),
		string(LocationConstraintItemLocationTypestreetAddress),
	}
}

func parseLocationConstraintItemLocationType(input string) (*LocationConstraintItemLocationType, error) {
	vals := map[string]LocationConstraintItemLocationType{
		"businessaddress": LocationConstraintItemLocationTypebusinessAddress,
		"conferenceroom":  LocationConstraintItemLocationTypeconferenceRoom,
		"default":         LocationConstraintItemLocationTypedefault,
		"geocoordinates":  LocationConstraintItemLocationTypegeoCoordinates,
		"homeaddress":     LocationConstraintItemLocationTypehomeAddress,
		"hotel":           LocationConstraintItemLocationTypehotel,
		"localbusiness":   LocationConstraintItemLocationTypelocalBusiness,
		"postaladdress":   LocationConstraintItemLocationTypepostalAddress,
		"restaurant":      LocationConstraintItemLocationTyperestaurant,
		"streetaddress":   LocationConstraintItemLocationTypestreetAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationConstraintItemLocationType(input)
	return &out, nil
}
