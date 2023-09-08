package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobilityManagementPolicyAppliesTo string

const (
	MobilityManagementPolicyAppliesToall      MobilityManagementPolicyAppliesTo = "All"
	MobilityManagementPolicyAppliesTonone     MobilityManagementPolicyAppliesTo = "None"
	MobilityManagementPolicyAppliesToselected MobilityManagementPolicyAppliesTo = "Selected"
)

func PossibleValuesForMobilityManagementPolicyAppliesTo() []string {
	return []string{
		string(MobilityManagementPolicyAppliesToall),
		string(MobilityManagementPolicyAppliesTonone),
		string(MobilityManagementPolicyAppliesToselected),
	}
}

func parseMobilityManagementPolicyAppliesTo(input string) (*MobilityManagementPolicyAppliesTo, error) {
	vals := map[string]MobilityManagementPolicyAppliesTo{
		"all":      MobilityManagementPolicyAppliesToall,
		"none":     MobilityManagementPolicyAppliesTonone,
		"selected": MobilityManagementPolicyAppliesToselected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobilityManagementPolicyAppliesTo(input)
	return &out, nil
}
