package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyAssignmentSource string

const (
	DeviceCompliancePolicyAssignmentSourcedirect     DeviceCompliancePolicyAssignmentSource = "Direct"
	DeviceCompliancePolicyAssignmentSourcepolicySets DeviceCompliancePolicyAssignmentSource = "PolicySets"
)

func PossibleValuesForDeviceCompliancePolicyAssignmentSource() []string {
	return []string{
		string(DeviceCompliancePolicyAssignmentSourcedirect),
		string(DeviceCompliancePolicyAssignmentSourcepolicySets),
	}
}

func parseDeviceCompliancePolicyAssignmentSource(input string) (*DeviceCompliancePolicyAssignmentSource, error) {
	vals := map[string]DeviceCompliancePolicyAssignmentSource{
		"direct":     DeviceCompliancePolicyAssignmentSourcedirect,
		"policysets": DeviceCompliancePolicyAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyAssignmentSource(input)
	return &out, nil
}
