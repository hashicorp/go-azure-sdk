package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedInboundDataTransferSources string

const (
	IosManagedAppProtectionAllowedInboundDataTransferSourcesallApps     IosManagedAppProtectionAllowedInboundDataTransferSources = "AllApps"
	IosManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps IosManagedAppProtectionAllowedInboundDataTransferSources = "ManagedApps"
	IosManagedAppProtectionAllowedInboundDataTransferSourcesnone        IosManagedAppProtectionAllowedInboundDataTransferSources = "None"
)

func PossibleValuesForIosManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(IosManagedAppProtectionAllowedInboundDataTransferSourcesallApps),
		string(IosManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps),
		string(IosManagedAppProtectionAllowedInboundDataTransferSourcesnone),
	}
}

func parseIosManagedAppProtectionAllowedInboundDataTransferSources(input string) (*IosManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]IosManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps":     IosManagedAppProtectionAllowedInboundDataTransferSourcesallApps,
		"managedapps": IosManagedAppProtectionAllowedInboundDataTransferSourcesmanagedApps,
		"none":        IosManagedAppProtectionAllowedInboundDataTransferSourcesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
