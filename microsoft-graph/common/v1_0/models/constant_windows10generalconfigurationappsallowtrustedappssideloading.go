package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading string

const (
	Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingallowed       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading = "Allowed"
	Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingblocked       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading = "Blocked"
	Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingnotConfigured Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationAppsAllowTrustedAppsSideloading() []string {
	return []string{
		string(Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingallowed),
		string(Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingblocked),
		string(Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingnotConfigured),
	}
}

func parseWindows10GeneralConfigurationAppsAllowTrustedAppsSideloading(input string) (*Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading, error) {
	vals := map[string]Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading{
		"allowed":       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingallowed,
		"blocked":       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingblocked,
		"notconfigured": Windows10GeneralConfigurationAppsAllowTrustedAppsSideloadingnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading(input)
	return &out, nil
}
