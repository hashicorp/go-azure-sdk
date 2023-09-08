package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionManagedBrowser string

const (
	IosManagedAppProtectionManagedBrowsermicrosoftEdge IosManagedAppProtectionManagedBrowser = "MicrosoftEdge"
	IosManagedAppProtectionManagedBrowsernotConfigured IosManagedAppProtectionManagedBrowser = "NotConfigured"
)

func PossibleValuesForIosManagedAppProtectionManagedBrowser() []string {
	return []string{
		string(IosManagedAppProtectionManagedBrowsermicrosoftEdge),
		string(IosManagedAppProtectionManagedBrowsernotConfigured),
	}
}

func parseIosManagedAppProtectionManagedBrowser(input string) (*IosManagedAppProtectionManagedBrowser, error) {
	vals := map[string]IosManagedAppProtectionManagedBrowser{
		"microsoftedge": IosManagedAppProtectionManagedBrowsermicrosoftEdge,
		"notconfigured": IosManagedAppProtectionManagedBrowsernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionManagedBrowser(input)
	return &out, nil
}
