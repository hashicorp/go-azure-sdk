package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskUWPAppAppType string

const (
	WindowsKioskUWPAppAppTypeaumId   WindowsKioskUWPAppAppType = "AumId"
	WindowsKioskUWPAppAppTypedesktop WindowsKioskUWPAppAppType = "Desktop"
	WindowsKioskUWPAppAppTypestore   WindowsKioskUWPAppAppType = "Store"
	WindowsKioskUWPAppAppTypeunknown WindowsKioskUWPAppAppType = "Unknown"
)

func PossibleValuesForWindowsKioskUWPAppAppType() []string {
	return []string{
		string(WindowsKioskUWPAppAppTypeaumId),
		string(WindowsKioskUWPAppAppTypedesktop),
		string(WindowsKioskUWPAppAppTypestore),
		string(WindowsKioskUWPAppAppTypeunknown),
	}
}

func parseWindowsKioskUWPAppAppType(input string) (*WindowsKioskUWPAppAppType, error) {
	vals := map[string]WindowsKioskUWPAppAppType{
		"aumid":   WindowsKioskUWPAppAppTypeaumId,
		"desktop": WindowsKioskUWPAppAppTypedesktop,
		"store":   WindowsKioskUWPAppAppTypestore,
		"unknown": WindowsKioskUWPAppAppTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskUWPAppAppType(input)
	return &out, nil
}
