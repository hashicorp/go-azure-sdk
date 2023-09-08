package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CountryNamedLocationCountryLookupMethod string

const (
	CountryNamedLocationCountryLookupMethodauthenticatorAppGps CountryNamedLocationCountryLookupMethod = "AuthenticatorAppGps"
	CountryNamedLocationCountryLookupMethodclientIpAddress     CountryNamedLocationCountryLookupMethod = "ClientIpAddress"
)

func PossibleValuesForCountryNamedLocationCountryLookupMethod() []string {
	return []string{
		string(CountryNamedLocationCountryLookupMethodauthenticatorAppGps),
		string(CountryNamedLocationCountryLookupMethodclientIpAddress),
	}
}

func parseCountryNamedLocationCountryLookupMethod(input string) (*CountryNamedLocationCountryLookupMethod, error) {
	vals := map[string]CountryNamedLocationCountryLookupMethod{
		"authenticatorappgps": CountryNamedLocationCountryLookupMethodauthenticatorAppGps,
		"clientipaddress":     CountryNamedLocationCountryLookupMethodclientIpAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CountryNamedLocationCountryLookupMethod(input)
	return &out, nil
}
