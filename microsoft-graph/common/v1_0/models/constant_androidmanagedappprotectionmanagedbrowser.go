package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionManagedBrowser string

const (
	AndroidManagedAppProtectionManagedBrowsermicrosoftEdge AndroidManagedAppProtectionManagedBrowser = "MicrosoftEdge"
	AndroidManagedAppProtectionManagedBrowsernotConfigured AndroidManagedAppProtectionManagedBrowser = "NotConfigured"
)

func PossibleValuesForAndroidManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(AndroidManagedAppProtectionManagedBrowsermicrosoftEdge),
		string(AndroidManagedAppProtectionManagedBrowsernotConfigured),
	}
}

func parseAndroidManagedAppProtectionManagedBrowser(input string) (*AndroidManagedAppProtectionManagedBrowser, error) {
	vals := map[string]AndroidManagedAppProtectionManagedBrowser{
		"microsoftedge": AndroidManagedAppProtectionManagedBrowsermicrosoftEdge,
		"notconfigured": AndroidManagedAppProtectionManagedBrowsernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
