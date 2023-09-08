package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionPinCharacterSet string

const (
	DefaultManagedAppProtectionPinCharacterSetalphanumericAndSymbol DefaultManagedAppProtectionPinCharacterSet = "AlphanumericAndSymbol"
	DefaultManagedAppProtectionPinCharacterSetnumeric               DefaultManagedAppProtectionPinCharacterSet = "Numeric"
)

func PossibleValuesForDefaultManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(DefaultManagedAppProtectionPinCharacterSetalphanumericAndSymbol),
		string(DefaultManagedAppProtectionPinCharacterSetnumeric),
	}
}

func parseDefaultManagedAppProtectionPinCharacterSet(input string) (*DefaultManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]DefaultManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": DefaultManagedAppProtectionPinCharacterSetalphanumericAndSymbol,
		"numeric":               DefaultManagedAppProtectionPinCharacterSetnumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
