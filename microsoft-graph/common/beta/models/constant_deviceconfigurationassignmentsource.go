package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationAssignmentSource string

const (
	DeviceConfigurationAssignmentSourcedirect     DeviceConfigurationAssignmentSource = "Direct"
	DeviceConfigurationAssignmentSourcepolicySets DeviceConfigurationAssignmentSource = "PolicySets"
)

func PossibleValuesForDeviceConfigurationAssignmentSource() []string {
	return []string{
		string(DeviceConfigurationAssignmentSourcedirect),
		string(DeviceConfigurationAssignmentSourcepolicySets),
	}
}

func parseDeviceConfigurationAssignmentSource(input string) (*DeviceConfigurationAssignmentSource, error) {
	vals := map[string]DeviceConfigurationAssignmentSource{
		"direct":     DeviceConfigurationAssignmentSourcedirect,
		"policysets": DeviceConfigurationAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationAssignmentSource(input)
	return &out, nil
}
