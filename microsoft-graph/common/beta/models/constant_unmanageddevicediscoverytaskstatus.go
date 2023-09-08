package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmanagedDeviceDiscoveryTaskStatus string

const (
	UnmanagedDeviceDiscoveryTaskStatusactive    UnmanagedDeviceDiscoveryTaskStatus = "Active"
	UnmanagedDeviceDiscoveryTaskStatuscompleted UnmanagedDeviceDiscoveryTaskStatus = "Completed"
	UnmanagedDeviceDiscoveryTaskStatuspending   UnmanagedDeviceDiscoveryTaskStatus = "Pending"
	UnmanagedDeviceDiscoveryTaskStatusrejected  UnmanagedDeviceDiscoveryTaskStatus = "Rejected"
	UnmanagedDeviceDiscoveryTaskStatusunknown   UnmanagedDeviceDiscoveryTaskStatus = "Unknown"
)

func PossibleValuesForUnmanagedDeviceDiscoveryTaskStatus() []string {
	return []string{
		string(UnmanagedDeviceDiscoveryTaskStatusactive),
		string(UnmanagedDeviceDiscoveryTaskStatuscompleted),
		string(UnmanagedDeviceDiscoveryTaskStatuspending),
		string(UnmanagedDeviceDiscoveryTaskStatusrejected),
		string(UnmanagedDeviceDiscoveryTaskStatusunknown),
	}
}

func parseUnmanagedDeviceDiscoveryTaskStatus(input string) (*UnmanagedDeviceDiscoveryTaskStatus, error) {
	vals := map[string]UnmanagedDeviceDiscoveryTaskStatus{
		"active":    UnmanagedDeviceDiscoveryTaskStatusactive,
		"completed": UnmanagedDeviceDiscoveryTaskStatuscompleted,
		"pending":   UnmanagedDeviceDiscoveryTaskStatuspending,
		"rejected":  UnmanagedDeviceDiscoveryTaskStatusrejected,
		"unknown":   UnmanagedDeviceDiscoveryTaskStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmanagedDeviceDiscoveryTaskStatus(input)
	return &out, nil
}
