package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDeviceDiscoveryTaskCategory string

const (
	UnmanagedDeviceDiscoveryTaskCategoryadvancedThreatProtection UnmanagedDeviceDiscoveryTaskCategory = "AdvancedThreatProtection"
	UnmanagedDeviceDiscoveryTaskCategoryunknown                  UnmanagedDeviceDiscoveryTaskCategory = "Unknown"
)

func PossibleValuesForUnmanagedDeviceDiscoveryTaskCategory() []string {
	return []string{
		string(UnmanagedDeviceDiscoveryTaskCategoryadvancedThreatProtection),
		string(UnmanagedDeviceDiscoveryTaskCategoryunknown),
	}
}

func parseUnmanagedDeviceDiscoveryTaskCategory(input string) (*UnmanagedDeviceDiscoveryTaskCategory, error) {
	vals := map[string]UnmanagedDeviceDiscoveryTaskCategory{
		"advancedthreatprotection": UnmanagedDeviceDiscoveryTaskCategoryadvancedThreatProtection,
		"unknown":                  UnmanagedDeviceDiscoveryTaskCategoryunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmanagedDeviceDiscoveryTaskCategory(input)
	return &out, nil
}
