package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionPinCharacterSet string

const (
	AndroidManagedAppProtectionPinCharacterSetalphanumericAndSymbol AndroidManagedAppProtectionPinCharacterSet = "AlphanumericAndSymbol"
	AndroidManagedAppProtectionPinCharacterSetnumeric               AndroidManagedAppProtectionPinCharacterSet = "Numeric"
)

func PossibleValuesForAndroidManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(AndroidManagedAppProtectionPinCharacterSetalphanumericAndSymbol),
		string(AndroidManagedAppProtectionPinCharacterSetnumeric),
	}
}

func parseAndroidManagedAppProtectionPinCharacterSet(input string) (*AndroidManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]AndroidManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": AndroidManagedAppProtectionPinCharacterSetalphanumericAndSymbol,
		"numeric":               AndroidManagedAppProtectionPinCharacterSetnumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
