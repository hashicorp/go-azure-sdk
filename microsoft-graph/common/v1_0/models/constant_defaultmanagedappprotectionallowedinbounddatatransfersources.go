package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedInboundDataTransferSources string

const (
	DefaultManagedAppProtectionAllowedInboundDataTransferSourcesallApps     DefaultManagedAppProtectionAllowedInboundDataTransferSources = "AllApps"
	DefaultManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps DefaultManagedAppProtectionAllowedInboundDataTransferSources = "ManagedApps"
	DefaultManagedAppProtectionAllowedInboundDataTransferSourcesnone        DefaultManagedAppProtectionAllowedInboundDataTransferSources = "None"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedInboundDataTransferSourcesallApps),
		string(DefaultManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps),
		string(DefaultManagedAppProtectionAllowedInboundDataTransferSourcesnone),
	}
}

func parseDefaultManagedAppProtectionAllowedInboundDataTransferSources(input string) (*DefaultManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     DefaultManagedAppProtectionAllowedInboundDataTransferSourcesallApps,
		"managedapps": DefaultManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps,
		"none":        DefaultManagedAppProtectionAllowedInboundDataTransferSourcesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
