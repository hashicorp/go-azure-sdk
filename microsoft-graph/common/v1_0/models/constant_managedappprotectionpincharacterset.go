package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionPinCharacterSet string

const (
	ManagedAppProtectionPinCharacterSetalphanumericAndSymbol ManagedAppProtectionPinCharacterSet = "AlphanumericAndSymbol"
	ManagedAppProtectionPinCharacterSetnumeric               ManagedAppProtectionPinCharacterSet = "Numeric"
)

func PossibleValuesForManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(ManagedAppProtectionPinCharacterSetalphanumericAndSymbol),
		string(ManagedAppProtectionPinCharacterSetnumeric),
	}
}

func parseManagedAppProtectionPinCharacterSet(input string) (*ManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]ManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": ManagedAppProtectionPinCharacterSetalphanumericAndSymbol,
		"numeric":               ManagedAppProtectionPinCharacterSetnumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
