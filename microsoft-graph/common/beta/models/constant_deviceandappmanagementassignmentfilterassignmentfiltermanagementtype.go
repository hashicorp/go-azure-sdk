package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType string

const (
	DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeapps    DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType = "Apps"
	DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypedevices DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType = "Devices"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeapps),
		string(DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypedevices),
	}
}

func parseDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(input string) (*DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType{
		"apps":    DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypeapps,
		"devices": DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementTypedevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(input)
	return &out, nil
}
