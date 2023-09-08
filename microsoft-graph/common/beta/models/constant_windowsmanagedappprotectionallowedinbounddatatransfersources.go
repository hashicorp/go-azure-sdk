package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionAllowedInboundDataTransferSources string

const (
	WindowsManagedAppProtectionAllowedInboundDataTransferSourcesallApps WindowsManagedAppProtectionAllowedInboundDataTransferSources = "AllApps"
	WindowsManagedAppProtectionAllowedInboundDataTransferSourcesnone    WindowsManagedAppProtectionAllowedInboundDataTransferSources = "None"
)

func PossibleValuesForWindowsManagedAppProtectionAllowedInboundDataTransferSources() []string {
	return []string{
		string(WindowsManagedAppProtectionAllowedInboundDataTransferSourcesallApps),
		string(WindowsManagedAppProtectionAllowedInboundDataTransferSourcesnone),
	}
}

func parseWindowsManagedAppProtectionAllowedInboundDataTransferSources(input string) (*WindowsManagedAppProtectionAllowedInboundDataTransferSources, error) {
	vals := map[string]WindowsManagedAppProtectionAllowedInboundDataTransferSources{
		"allapps": WindowsManagedAppProtectionAllowedInboundDataTransferSourcesallApps,
		"none":    WindowsManagedAppProtectionAllowedInboundDataTransferSourcesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionAllowedInboundDataTransferSources(input)
	return &out, nil
}
