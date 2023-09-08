package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationConstraintItemUniqueIdType string

const (
	LocationConstraintItemUniqueIdTypebing          LocationConstraintItemUniqueIdType = "Bing"
	LocationConstraintItemUniqueIdTypedirectory     LocationConstraintItemUniqueIdType = "Directory"
	LocationConstraintItemUniqueIdTypelocationStore LocationConstraintItemUniqueIdType = "LocationStore"
	LocationConstraintItemUniqueIdTypeprivate       LocationConstraintItemUniqueIdType = "Private"
	LocationConstraintItemUniqueIdTypeunknown       LocationConstraintItemUniqueIdType = "Unknown"
)

func PossibleValuesForLocationConstraintItemUniqueIdType() []string {
	return []string{
		string(LocationConstraintItemUniqueIdTypebing),
		string(LocationConstraintItemUniqueIdTypedirectory),
		string(LocationConstraintItemUniqueIdTypelocationStore),
		string(LocationConstraintItemUniqueIdTypeprivate),
		string(LocationConstraintItemUniqueIdTypeunknown),
	}
}

func parseLocationConstraintItemUniqueIdType(input string) (*LocationConstraintItemUniqueIdType, error) {
	vals := map[string]LocationConstraintItemUniqueIdType{
		"bing":          LocationConstraintItemUniqueIdTypebing,
		"directory":     LocationConstraintItemUniqueIdTypedirectory,
		"locationstore": LocationConstraintItemUniqueIdTypelocationStore,
		"private":       LocationConstraintItemUniqueIdTypeprivate,
		"unknown":       LocationConstraintItemUniqueIdTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocationConstraintItemUniqueIdType(input)
	return &out, nil
}
