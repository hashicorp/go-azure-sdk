package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionPinCharacterSet string

const (
	TargetedManagedAppProtectionPinCharacterSetalphanumericAndSymbol TargetedManagedAppProtectionPinCharacterSet = "AlphanumericAndSymbol"
	TargetedManagedAppProtectionPinCharacterSetnumeric               TargetedManagedAppProtectionPinCharacterSet = "Numeric"
)

func PossibleValuesForTargetedManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(TargetedManagedAppProtectionPinCharacterSetalphanumericAndSymbol),
		string(TargetedManagedAppProtectionPinCharacterSetnumeric),
	}
}

func parseTargetedManagedAppProtectionPinCharacterSet(input string) (*TargetedManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]TargetedManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": TargetedManagedAppProtectionPinCharacterSetalphanumericAndSymbol,
		"numeric":               TargetedManagedAppProtectionPinCharacterSetnumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
