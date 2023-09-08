package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppPolicyAssignmentSource string

const (
	TargetedManagedAppPolicyAssignmentSourcedirect     TargetedManagedAppPolicyAssignmentSource = "Direct"
	TargetedManagedAppPolicyAssignmentSourcepolicySets TargetedManagedAppPolicyAssignmentSource = "PolicySets"
)

func PossibleValuesForTargetedManagedAppPolicyAssignmentSource() []string {
	return []string{
		string(TargetedManagedAppPolicyAssignmentSourcedirect),
		string(TargetedManagedAppPolicyAssignmentSourcepolicySets),
	}
}

func parseTargetedManagedAppPolicyAssignmentSource(input string) (*TargetedManagedAppPolicyAssignmentSource, error) {
	vals := map[string]TargetedManagedAppPolicyAssignmentSource{
		"direct":     TargetedManagedAppPolicyAssignmentSourcedirect,
		"policysets": TargetedManagedAppPolicyAssignmentSourcepolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppPolicyAssignmentSource(input)
	return &out, nil
}
