package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionPolicySetItemStatus string

const (
	ManagedAppProtectionPolicySetItemStatuserror          ManagedAppProtectionPolicySetItemStatus = "Error"
	ManagedAppProtectionPolicySetItemStatusnotAssigned    ManagedAppProtectionPolicySetItemStatus = "NotAssigned"
	ManagedAppProtectionPolicySetItemStatuspartialSuccess ManagedAppProtectionPolicySetItemStatus = "PartialSuccess"
	ManagedAppProtectionPolicySetItemStatussuccess        ManagedAppProtectionPolicySetItemStatus = "Success"
	ManagedAppProtectionPolicySetItemStatusunknown        ManagedAppProtectionPolicySetItemStatus = "Unknown"
	ManagedAppProtectionPolicySetItemStatusvalidating     ManagedAppProtectionPolicySetItemStatus = "Validating"
)

func PossibleValuesForManagedAppProtectionPolicySetItemStatus() []string {
	return []string{
		string(ManagedAppProtectionPolicySetItemStatuserror),
		string(ManagedAppProtectionPolicySetItemStatusnotAssigned),
		string(ManagedAppProtectionPolicySetItemStatuspartialSuccess),
		string(ManagedAppProtectionPolicySetItemStatussuccess),
		string(ManagedAppProtectionPolicySetItemStatusunknown),
		string(ManagedAppProtectionPolicySetItemStatusvalidating),
	}
}

func parseManagedAppProtectionPolicySetItemStatus(input string) (*ManagedAppProtectionPolicySetItemStatus, error) {
	vals := map[string]ManagedAppProtectionPolicySetItemStatus{
		"error":          ManagedAppProtectionPolicySetItemStatuserror,
		"notassigned":    ManagedAppProtectionPolicySetItemStatusnotAssigned,
		"partialsuccess": ManagedAppProtectionPolicySetItemStatuspartialSuccess,
		"success":        ManagedAppProtectionPolicySetItemStatussuccess,
		"unknown":        ManagedAppProtectionPolicySetItemStatusunknown,
		"validating":     ManagedAppProtectionPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionPolicySetItemStatus(input)
	return &out, nil
}
