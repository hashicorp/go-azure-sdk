package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppConfigurationPolicySetItemErrorCode string

const (
	TargetedManagedAppConfigurationPolicySetItemErrorCodedeleted      TargetedManagedAppConfigurationPolicySetItemErrorCode = "Deleted"
	TargetedManagedAppConfigurationPolicySetItemErrorCodenoError      TargetedManagedAppConfigurationPolicySetItemErrorCode = "NoError"
	TargetedManagedAppConfigurationPolicySetItemErrorCodenotFound     TargetedManagedAppConfigurationPolicySetItemErrorCode = "NotFound"
	TargetedManagedAppConfigurationPolicySetItemErrorCodeunauthorized TargetedManagedAppConfigurationPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForTargetedManagedAppConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(TargetedManagedAppConfigurationPolicySetItemErrorCodedeleted),
		string(TargetedManagedAppConfigurationPolicySetItemErrorCodenoError),
		string(TargetedManagedAppConfigurationPolicySetItemErrorCodenotFound),
		string(TargetedManagedAppConfigurationPolicySetItemErrorCodeunauthorized),
	}
}

func parseTargetedManagedAppConfigurationPolicySetItemErrorCode(input string) (*TargetedManagedAppConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]TargetedManagedAppConfigurationPolicySetItemErrorCode{
		"deleted":      TargetedManagedAppConfigurationPolicySetItemErrorCodedeleted,
		"noerror":      TargetedManagedAppConfigurationPolicySetItemErrorCodenoError,
		"notfound":     TargetedManagedAppConfigurationPolicySetItemErrorCodenotFound,
		"unauthorized": TargetedManagedAppConfigurationPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
