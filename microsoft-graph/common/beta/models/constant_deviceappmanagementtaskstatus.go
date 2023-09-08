package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTaskStatus string

const (
	DeviceAppManagementTaskStatusactive    DeviceAppManagementTaskStatus = "Active"
	DeviceAppManagementTaskStatuscompleted DeviceAppManagementTaskStatus = "Completed"
	DeviceAppManagementTaskStatuspending   DeviceAppManagementTaskStatus = "Pending"
	DeviceAppManagementTaskStatusrejected  DeviceAppManagementTaskStatus = "Rejected"
	DeviceAppManagementTaskStatusunknown   DeviceAppManagementTaskStatus = "Unknown"
)

func PossibleValuesForDeviceAppManagementTaskStatus() []string {
	return []string{
		string(DeviceAppManagementTaskStatusactive),
		string(DeviceAppManagementTaskStatuscompleted),
		string(DeviceAppManagementTaskStatuspending),
		string(DeviceAppManagementTaskStatusrejected),
		string(DeviceAppManagementTaskStatusunknown),
	}
}

func parseDeviceAppManagementTaskStatus(input string) (*DeviceAppManagementTaskStatus, error) {
	vals := map[string]DeviceAppManagementTaskStatus{
		"active":    DeviceAppManagementTaskStatusactive,
		"completed": DeviceAppManagementTaskStatuscompleted,
		"pending":   DeviceAppManagementTaskStatuspending,
		"rejected":  DeviceAppManagementTaskStatusrejected,
		"unknown":   DeviceAppManagementTaskStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementTaskStatus(input)
	return &out, nil
}
