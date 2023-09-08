package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagedDeviceOwnerType string

const (
	WindowsManagedDeviceManagedDeviceOwnerTypecompany  WindowsManagedDeviceManagedDeviceOwnerType = "Company"
	WindowsManagedDeviceManagedDeviceOwnerTypepersonal WindowsManagedDeviceManagedDeviceOwnerType = "Personal"
	WindowsManagedDeviceManagedDeviceOwnerTypeunknown  WindowsManagedDeviceManagedDeviceOwnerType = "Unknown"
)

func PossibleValuesForWindowsManagedDeviceManagedDeviceOwnerType() []string {
	return []string{
		string(WindowsManagedDeviceManagedDeviceOwnerTypecompany),
		string(WindowsManagedDeviceManagedDeviceOwnerTypepersonal),
		string(WindowsManagedDeviceManagedDeviceOwnerTypeunknown),
	}
}

func parseWindowsManagedDeviceManagedDeviceOwnerType(input string) (*WindowsManagedDeviceManagedDeviceOwnerType, error) {
	vals := map[string]WindowsManagedDeviceManagedDeviceOwnerType{
		"company":  WindowsManagedDeviceManagedDeviceOwnerTypecompany,
		"personal": WindowsManagedDeviceManagedDeviceOwnerTypepersonal,
		"unknown":  WindowsManagedDeviceManagedDeviceOwnerTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagedDeviceOwnerType(input)
	return &out, nil
}
