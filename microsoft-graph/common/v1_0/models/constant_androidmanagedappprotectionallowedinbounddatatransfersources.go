package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedInboundDataTransferSources string

const (
	AndroidManagedAppProtectionAllowedInboundDataTransferSourcesallApps     AndroidManagedAppProtectionAllowedInboundDataTransferSources = "AllApps"
	AndroidManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps AndroidManagedAppProtectionAllowedInboundDataTransferSources = "ManagedApps"
	AndroidManagedAppProtectionAllowedInboundDataTransferSourcesnone        AndroidManagedAppProtectionAllowedInboundDataTransferSources = "None"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedInboundDataTransferSourcesallApps),
		string(AndroidManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps),
		string(AndroidManagedAppProtectionAllowedInboundDataTransferSourcesnone),
	}
}

func parseAndroidManagedAppProtectionAllowedInboundDataTransferSources(input string) (*AndroidManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     AndroidManagedAppProtectionAllowedInboundDataTransferSourcesallApps,
		"managedapps": AndroidManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps,
		"none":        AndroidManagedAppProtectionAllowedInboundDataTransferSourcesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
