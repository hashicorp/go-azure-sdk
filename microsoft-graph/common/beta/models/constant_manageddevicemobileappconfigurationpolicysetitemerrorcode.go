package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode string

const (
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodedeleted      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "Deleted"
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodenoError      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "NoError"
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodenotFound     ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "NotFound"
	ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodeunauthorized ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodedeleted),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodenoError),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodenotFound),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodeunauthorized),
	}
}

func parseManagedDeviceMobileAppConfigurationPolicySetItemErrorCode(input string) (*ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode{
		"deleted":      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodedeleted,
		"noerror":      ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodenoError,
		"notfound":     ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodenotFound,
		"unauthorized": ManagedDeviceMobileAppConfigurationPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
