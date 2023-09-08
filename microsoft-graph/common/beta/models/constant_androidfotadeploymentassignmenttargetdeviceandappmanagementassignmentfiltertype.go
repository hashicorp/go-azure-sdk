package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Exclude"
	AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Include"
	AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone    AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "None"
)

func PossibleValuesForAndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude),
		string(AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude),
		string(AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone),
	}
}

func parseAndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude,
		"include": AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude,
		"none":    AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidFotaDeploymentAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
