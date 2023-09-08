package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Exclude"
	AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Include"
	AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone    AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "None"
)

func PossibleValuesForAllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude),
		string(AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude),
		string(AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone),
	}
}

func parseAllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude,
		"include": AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude,
		"none":    AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
