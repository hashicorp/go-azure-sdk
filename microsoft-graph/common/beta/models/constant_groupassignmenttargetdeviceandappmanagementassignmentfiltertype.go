package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Exclude"
	GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Include"
	GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone    GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "None"
)

func PossibleValuesForGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude),
		string(GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude),
		string(GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone),
	}
}

func parseGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude,
		"include": GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude,
		"none":    GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
