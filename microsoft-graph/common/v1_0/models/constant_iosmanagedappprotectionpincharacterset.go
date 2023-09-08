package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionPinCharacterSet string

const (
	IosManagedAppProtectionPinCharacterSetalphanumericAndSymbol IosManagedAppProtectionPinCharacterSet = "AlphanumericAndSymbol"
	IosManagedAppProtectionPinCharacterSetnumeric               IosManagedAppProtectionPinCharacterSet = "Numeric"
)

func PossibleValuesForIosManagedAppProtectionPinCharacterSet() []string {
	return []string{
		string(IosManagedAppProtectionPinCharacterSetalphanumericAndSymbol),
		string(IosManagedAppProtectionPinCharacterSetnumeric),
	}
}

func parseIosManagedAppProtectionPinCharacterSet(input string) (*IosManagedAppProtectionPinCharacterSet, error) {
	vals := map[string]IosManagedAppProtectionPinCharacterSet{
		"alphanumericandsymbol": IosManagedAppProtectionPinCharacterSetalphanumericAndSymbol,
		"numeric":               IosManagedAppProtectionPinCharacterSetnumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionPinCharacterSet(input)
	return &out, nil
}
