package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationPolicySetItemStatus string

const (
	TargetedManagedAppConfigurationPolicySetItemStatuserror          TargetedManagedAppConfigurationPolicySetItemStatus = "Error"
	TargetedManagedAppConfigurationPolicySetItemStatusnotAssigned    TargetedManagedAppConfigurationPolicySetItemStatus = "NotAssigned"
	TargetedManagedAppConfigurationPolicySetItemStatuspartialSuccess TargetedManagedAppConfigurationPolicySetItemStatus = "PartialSuccess"
	TargetedManagedAppConfigurationPolicySetItemStatussuccess        TargetedManagedAppConfigurationPolicySetItemStatus = "Success"
	TargetedManagedAppConfigurationPolicySetItemStatusunknown        TargetedManagedAppConfigurationPolicySetItemStatus = "Unknown"
	TargetedManagedAppConfigurationPolicySetItemStatusvalidating     TargetedManagedAppConfigurationPolicySetItemStatus = "Validating"
)

func PossibleValuesForTargetedManagedAppConfigurationPolicySetItemStatus() []string {
	return []string{
		string(TargetedManagedAppConfigurationPolicySetItemStatuserror),
		string(TargetedManagedAppConfigurationPolicySetItemStatusnotAssigned),
		string(TargetedManagedAppConfigurationPolicySetItemStatuspartialSuccess),
		string(TargetedManagedAppConfigurationPolicySetItemStatussuccess),
		string(TargetedManagedAppConfigurationPolicySetItemStatusunknown),
		string(TargetedManagedAppConfigurationPolicySetItemStatusvalidating),
	}
}

func parseTargetedManagedAppConfigurationPolicySetItemStatus(input string) (*TargetedManagedAppConfigurationPolicySetItemStatus, error) {
	vals := map[string]TargetedManagedAppConfigurationPolicySetItemStatus{
		"error":          TargetedManagedAppConfigurationPolicySetItemStatuserror,
		"notassigned":    TargetedManagedAppConfigurationPolicySetItemStatusnotAssigned,
		"partialsuccess": TargetedManagedAppConfigurationPolicySetItemStatuspartialSuccess,
		"success":        TargetedManagedAppConfigurationPolicySetItemStatussuccess,
		"unknown":        TargetedManagedAppConfigurationPolicySetItemStatusunknown,
		"validating":     TargetedManagedAppConfigurationPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationPolicySetItemStatus(input)
	return &out, nil
}
