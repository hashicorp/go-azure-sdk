package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationPolicySetItemStatus string

const (
	ManagedDeviceMobileAppConfigurationPolicySetItemStatuserror          ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "Error"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatusnotAssigned    ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "NotAssigned"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatuspartialSuccess ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "PartialSuccess"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatussuccess        ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "Success"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatusunknown        ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "Unknown"
	ManagedDeviceMobileAppConfigurationPolicySetItemStatusvalidating     ManagedDeviceMobileAppConfigurationPolicySetItemStatus = "Validating"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationPolicySetItemStatus() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatuserror),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatusnotAssigned),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatuspartialSuccess),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatussuccess),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatusunknown),
		string(ManagedDeviceMobileAppConfigurationPolicySetItemStatusvalidating),
	}
}

func parseManagedDeviceMobileAppConfigurationPolicySetItemStatus(input string) (*ManagedDeviceMobileAppConfigurationPolicySetItemStatus, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationPolicySetItemStatus{
		"error":          ManagedDeviceMobileAppConfigurationPolicySetItemStatuserror,
		"notassigned":    ManagedDeviceMobileAppConfigurationPolicySetItemStatusnotAssigned,
		"partialsuccess": ManagedDeviceMobileAppConfigurationPolicySetItemStatuspartialSuccess,
		"success":        ManagedDeviceMobileAppConfigurationPolicySetItemStatussuccess,
		"unknown":        ManagedDeviceMobileAppConfigurationPolicySetItemStatusunknown,
		"validating":     ManagedDeviceMobileAppConfigurationPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationPolicySetItemStatus(input)
	return &out, nil
}
