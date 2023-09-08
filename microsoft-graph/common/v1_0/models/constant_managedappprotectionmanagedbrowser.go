package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionManagedBrowser string

const (
	ManagedAppProtectionManagedBrowsermicrosoftEdge ManagedAppProtectionManagedBrowser = "MicrosoftEdge"
	ManagedAppProtectionManagedBrowsernotConfigured ManagedAppProtectionManagedBrowser = "NotConfigured"
)

func PossibleValuesForManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(ManagedAppProtectionManagedBrowsermicrosoftEdge),
		string(ManagedAppProtectionManagedBrowsernotConfigured),
	}
}

func parseManagedAppProtectionManagedBrowser(input string) (*ManagedAppProtectionManagedBrowser, error) {
	vals := map[string]ManagedAppProtectionManagedBrowser{
		"microsoftedge": ManagedAppProtectionManagedBrowsermicrosoftEdge,
		"notconfigured": ManagedAppProtectionManagedBrowsernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
