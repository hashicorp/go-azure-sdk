package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Exclude"
	ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Include"
	ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone    ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "None"
)

func PossibleValuesForExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude),
		string(ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude),
		string(ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone),
	}
}

func parseExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude,
		"include": ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude,
		"none":    ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
