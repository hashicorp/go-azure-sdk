package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceOwnerType string

const (
	ManagedDeviceOwnerTypecompany  ManagedDeviceOwnerType = "Company"
	ManagedDeviceOwnerTypepersonal ManagedDeviceOwnerType = "Personal"
	ManagedDeviceOwnerTypeunknown  ManagedDeviceOwnerType = "Unknown"
)

func PossibleValuesForManagedDeviceOwnerType() []string {
	return []string{
		string(ManagedDeviceOwnerTypecompany),
		string(ManagedDeviceOwnerTypepersonal),
		string(ManagedDeviceOwnerTypeunknown),
	}
}

func parseManagedDeviceOwnerType(input string) (*ManagedDeviceOwnerType, error) {
	vals := map[string]ManagedDeviceOwnerType{
		"company":  ManagedDeviceOwnerTypecompany,
		"personal": ManagedDeviceOwnerTypepersonal,
		"unknown":  ManagedDeviceOwnerTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceOwnerType(input)
	return &out, nil
}
