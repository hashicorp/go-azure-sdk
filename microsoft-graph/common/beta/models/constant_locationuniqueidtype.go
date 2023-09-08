package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationUniqueIdType string

const (
	LocationUniqueIdTypebing          LocationUniqueIdType = "Bing"
	LocationUniqueIdTypedirectory     LocationUniqueIdType = "Directory"
	LocationUniqueIdTypelocationStore LocationUniqueIdType = "LocationStore"
	LocationUniqueIdTypeprivate       LocationUniqueIdType = "Private"
	LocationUniqueIdTypeunknown       LocationUniqueIdType = "Unknown"
)

func PossibleValuesForLocationUniqueIdType() []string {
	return []string{
		string(LocationUniqueIdTypebing),
		string(LocationUniqueIdTypedirectory),
		string(LocationUniqueIdTypelocationStore),
		string(LocationUniqueIdTypeprivate),
		string(LocationUniqueIdTypeunknown),
	}
}

func parseLocationUniqueIdType(input string) (*LocationUniqueIdType, error) {
	vals := map[string]LocationUniqueIdType{
		"bing":          LocationUniqueIdTypebing,
		"directory":     LocationUniqueIdTypedirectory,
		"locationstore": LocationUniqueIdTypelocationStore,
		"private":       LocationUniqueIdTypeprivate,
		"unknown":       LocationUniqueIdTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationUniqueIdType(input)
	return &out, nil
}
