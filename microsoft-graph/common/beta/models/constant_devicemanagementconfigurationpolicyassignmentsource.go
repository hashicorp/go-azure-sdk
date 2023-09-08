package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyAssignmentSource string

const (
	DeviceManagementConfigurationPolicyAssignmentSourcedirect     DeviceManagementConfigurationPolicyAssignmentSource = "Direct"
	DeviceManagementConfigurationPolicyAssignmentSourcepolicySets DeviceManagementConfigurationPolicyAssignmentSource = "PolicySets"
)

func PossibleValuesForDeviceManagementConfigurationPolicyAssignmentSource() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyAssignmentSourcedirect),
		string(DeviceManagementConfigurationPolicyAssignmentSourcepolicySets),
	}
}

func parseDeviceManagementConfigurationPolicyAssignmentSource(input string) (*DeviceManagementConfigurationPolicyAssignmentSource, error) {
	vals := map[string]DeviceManagementConfigurationPolicyAssignmentSource{
		"direct":     DeviceManagementConfigurationPolicyAssignmentSourcedirect,
		"policysets": DeviceManagementConfigurationPolicyAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyAssignmentSource(input)
	return &out, nil
}
