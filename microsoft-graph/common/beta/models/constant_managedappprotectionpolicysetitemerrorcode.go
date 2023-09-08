package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionPolicySetItemErrorCode string

const (
	ManagedAppProtectionPolicySetItemErrorCodedeleted      ManagedAppProtectionPolicySetItemErrorCode = "Deleted"
	ManagedAppProtectionPolicySetItemErrorCodenoError      ManagedAppProtectionPolicySetItemErrorCode = "NoError"
	ManagedAppProtectionPolicySetItemErrorCodenotFound     ManagedAppProtectionPolicySetItemErrorCode = "NotFound"
	ManagedAppProtectionPolicySetItemErrorCodeunauthorized ManagedAppProtectionPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForManagedAppProtectionPolicySetItemErrorCode() []string {
	return []string{
		string(ManagedAppProtectionPolicySetItemErrorCodedeleted),
		string(ManagedAppProtectionPolicySetItemErrorCodenoError),
		string(ManagedAppProtectionPolicySetItemErrorCodenotFound),
		string(ManagedAppProtectionPolicySetItemErrorCodeunauthorized),
	}
}

func parseManagedAppProtectionPolicySetItemErrorCode(input string) (*ManagedAppProtectionPolicySetItemErrorCode, error) {
	vals := map[string]ManagedAppProtectionPolicySetItemErrorCode{
		"deleted":      ManagedAppProtectionPolicySetItemErrorCodedeleted,
		"noerror":      ManagedAppProtectionPolicySetItemErrorCodenoError,
		"notfound":     ManagedAppProtectionPolicySetItemErrorCodenotFound,
		"unauthorized": ManagedAppProtectionPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionPolicySetItemErrorCode(input)
	return &out, nil
}
