package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceOwnerType string

const (
	ComanagementEligibleDeviceOwnerTypecompany  ComanagementEligibleDeviceOwnerType = "Company"
	ComanagementEligibleDeviceOwnerTypepersonal ComanagementEligibleDeviceOwnerType = "Personal"
	ComanagementEligibleDeviceOwnerTypeunknown  ComanagementEligibleDeviceOwnerType = "Unknown"
)

func PossibleValuesForComanagementEligibleDeviceOwnerType() []string {
	return []string{
		string(ComanagementEligibleDeviceOwnerTypecompany),
		string(ComanagementEligibleDeviceOwnerTypepersonal),
		string(ComanagementEligibleDeviceOwnerTypeunknown),
	}
}

func parseComanagementEligibleDeviceOwnerType(input string) (*ComanagementEligibleDeviceOwnerType, error) {
	vals := map[string]ComanagementEligibleDeviceOwnerType{
		"company":  ComanagementEligibleDeviceOwnerTypecompany,
		"personal": ComanagementEligibleDeviceOwnerTypepersonal,
		"unknown":  ComanagementEligibleDeviceOwnerTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceOwnerType(input)
	return &out, nil
}
