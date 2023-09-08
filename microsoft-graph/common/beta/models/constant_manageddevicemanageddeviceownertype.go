package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagedDeviceOwnerType string

const (
	ManagedDeviceManagedDeviceOwnerTypecompany  ManagedDeviceManagedDeviceOwnerType = "Company"
	ManagedDeviceManagedDeviceOwnerTypepersonal ManagedDeviceManagedDeviceOwnerType = "Personal"
	ManagedDeviceManagedDeviceOwnerTypeunknown  ManagedDeviceManagedDeviceOwnerType = "Unknown"
)

func PossibleValuesForManagedDeviceManagedDeviceOwnerType() []string {
	return []string{
		string(ManagedDeviceManagedDeviceOwnerTypecompany),
		string(ManagedDeviceManagedDeviceOwnerTypepersonal),
		string(ManagedDeviceManagedDeviceOwnerTypeunknown),
	}
}

func parseManagedDeviceManagedDeviceOwnerType(input string) (*ManagedDeviceManagedDeviceOwnerType, error) {
	vals := map[string]ManagedDeviceManagedDeviceOwnerType{
		"company":  ManagedDeviceManagedDeviceOwnerTypecompany,
		"personal": ManagedDeviceManagedDeviceOwnerTypepersonal,
		"unknown":  ManagedDeviceManagedDeviceOwnerTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagedDeviceOwnerType(input)
	return &out, nil
}
