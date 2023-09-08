package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedInboundDataTransferSources string

const (
	ManagedAppProtectionAllowedInboundDataTransferSourcesallApps     ManagedAppProtectionAllowedInboundDataTransferSources = "AllApps"
	ManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps ManagedAppProtectionAllowedInboundDataTransferSources = "ManagedApps"
	ManagedAppProtectionAllowedInboundDataTransferSourcesnone        ManagedAppProtectionAllowedInboundDataTransferSources = "None"
)

func PossibleValuesForManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(ManagedAppProtectionAllowedInboundDataTransferSourcesallApps),
		string(ManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps),
		string(ManagedAppProtectionAllowedInboundDataTransferSourcesnone),
	}
}

func parseManagedAppProtectionAllowedInboundDataTransferSources(input string) (*ManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]ManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     ManagedAppProtectionAllowedInboundDataTransferSourcesallApps,
		"managedapps": ManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps,
		"none":        ManagedAppProtectionAllowedInboundDataTransferSourcesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
