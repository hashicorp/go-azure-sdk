package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppBaseAppType string

const (
	WindowsKioskAppBaseAppTypeaumId   WindowsKioskAppBaseAppType = "AumId"
	WindowsKioskAppBaseAppTypedesktop WindowsKioskAppBaseAppType = "Desktop"
	WindowsKioskAppBaseAppTypestore   WindowsKioskAppBaseAppType = "Store"
	WindowsKioskAppBaseAppTypeunknown WindowsKioskAppBaseAppType = "Unknown"
)

func PossibleValuesForWindowsKioskAppBaseAppType() []string {
	return []string{
		string(WindowsKioskAppBaseAppTypeaumId),
		string(WindowsKioskAppBaseAppTypedesktop),
		string(WindowsKioskAppBaseAppTypestore),
		string(WindowsKioskAppBaseAppTypeunknown),
	}
}

func parseWindowsKioskAppBaseAppType(input string) (*WindowsKioskAppBaseAppType, error) {
	vals := map[string]WindowsKioskAppBaseAppType{
		"aumid":   WindowsKioskAppBaseAppTypeaumId,
		"desktop": WindowsKioskAppBaseAppTypedesktop,
		"store":   WindowsKioskAppBaseAppTypestore,
		"unknown": WindowsKioskAppBaseAppTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskAppBaseAppType(input)
	return &out, nil
}
