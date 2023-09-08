package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Exclude"
	ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "Include"
	ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone    ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "None"
)

func PossibleValuesForConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude),
		string(ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude),
		string(ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone),
	}
}

func parseConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeexclude,
		"include": ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeinclude,
		"none":    ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerCollectionAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
