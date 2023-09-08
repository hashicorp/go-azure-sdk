package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDeviceDiscoveryTaskPriority string

const (
	UnmanagedDeviceDiscoveryTaskPriorityhigh UnmanagedDeviceDiscoveryTaskPriority = "High"
	UnmanagedDeviceDiscoveryTaskPrioritylow  UnmanagedDeviceDiscoveryTaskPriority = "Low"
	UnmanagedDeviceDiscoveryTaskPrioritynone UnmanagedDeviceDiscoveryTaskPriority = "None"
)

func PossibleValuesForUnmanagedDeviceDiscoveryTaskPriority() []string {
	return []string{
		string(UnmanagedDeviceDiscoveryTaskPriorityhigh),
		string(UnmanagedDeviceDiscoveryTaskPrioritylow),
		string(UnmanagedDeviceDiscoveryTaskPrioritynone),
	}
}

func parseUnmanagedDeviceDiscoveryTaskPriority(input string) (*UnmanagedDeviceDiscoveryTaskPriority, error) {
	vals := map[string]UnmanagedDeviceDiscoveryTaskPriority{
		"high": UnmanagedDeviceDiscoveryTaskPriorityhigh,
		"low":  UnmanagedDeviceDiscoveryTaskPrioritylow,
		"none": UnmanagedDeviceDiscoveryTaskPrioritynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmanagedDeviceDiscoveryTaskPriority(input)
	return &out, nil
}
