package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Exclude"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Include"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "None"
)

func PossibleValuesForDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone),
	}
}

func parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude,
		"include": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude,
		"none":    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
