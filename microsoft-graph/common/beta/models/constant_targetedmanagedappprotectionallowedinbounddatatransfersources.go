package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedInboundDataTransferSources string

const (
	TargetedManagedAppProtectionAllowedInboundDataTransferSourcesallApps     TargetedManagedAppProtectionAllowedInboundDataTransferSources = "AllApps"
	TargetedManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps TargetedManagedAppProtectionAllowedInboundDataTransferSources = "ManagedApps"
	TargetedManagedAppProtectionAllowedInboundDataTransferSourcesnone        TargetedManagedAppProtectionAllowedInboundDataTransferSources = "None"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedInboundDataTransferSourcesallApps),
		string(TargetedManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps),
		string(TargetedManagedAppProtectionAllowedInboundDataTransferSourcesnone),
	}
}

func parseTargetedManagedAppProtectionAllowedInboundDataTransferSources(input string) (*TargetedManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     TargetedManagedAppProtectionAllowedInboundDataTransferSourcesallApps,
		"managedapps": TargetedManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps,
		"none":        TargetedManagedAppProtectionAllowedInboundDataTransferSourcesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
