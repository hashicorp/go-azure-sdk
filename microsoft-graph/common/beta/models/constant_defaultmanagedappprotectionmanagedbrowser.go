package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionManagedBrowser string

const (
	DefaultManagedAppProtectionManagedBrowsermicrosoftEdge DefaultManagedAppProtectionManagedBrowser = "MicrosoftEdge"
	DefaultManagedAppProtectionManagedBrowsernotConfigured DefaultManagedAppProtectionManagedBrowser = "NotConfigured"
)

func PossibleValuesForDefaultManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(DefaultManagedAppProtectionManagedBrowsermicrosoftEdge),
		string(DefaultManagedAppProtectionManagedBrowsernotConfigured),
	}
}

func parseDefaultManagedAppProtectionManagedBrowser(input string) (*DefaultManagedAppProtectionManagedBrowser, error) {
	vals := map[string]DefaultManagedAppProtectionManagedBrowser{
		"microsoftedge": DefaultManagedAppProtectionManagedBrowsermicrosoftEdge,
		"notconfigured": DefaultManagedAppProtectionManagedBrowsernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
