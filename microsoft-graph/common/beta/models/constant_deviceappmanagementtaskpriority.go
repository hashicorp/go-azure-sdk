package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTaskPriority string

const (
	DeviceAppManagementTaskPriorityhigh DeviceAppManagementTaskPriority = "High"
	DeviceAppManagementTaskPrioritylow  DeviceAppManagementTaskPriority = "Low"
	DeviceAppManagementTaskPrioritynone DeviceAppManagementTaskPriority = "None"
)

func PossibleValuesForDeviceAppManagementTaskPriority() []string {
	return []string{
		string(DeviceAppManagementTaskPriorityhigh),
		string(DeviceAppManagementTaskPrioritylow),
		string(DeviceAppManagementTaskPrioritynone),
	}
}

func parseDeviceAppManagementTaskPriority(input string) (*DeviceAppManagementTaskPriority, error) {
	vals := map[string]DeviceAppManagementTaskPriority{
		"high": DeviceAppManagementTaskPriorityhigh,
		"low":  DeviceAppManagementTaskPrioritylow,
		"none": DeviceAppManagementTaskPrioritynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementTaskPriority(input)
	return &out, nil
}
