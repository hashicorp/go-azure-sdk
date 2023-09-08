package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationInkWorkspaceAccess string

const (
	Windows10GeneralConfigurationInkWorkspaceAccessdisabled      Windows10GeneralConfigurationInkWorkspaceAccess = "Disabled"
	Windows10GeneralConfigurationInkWorkspaceAccessenabled       Windows10GeneralConfigurationInkWorkspaceAccess = "Enabled"
	Windows10GeneralConfigurationInkWorkspaceAccessnotConfigured Windows10GeneralConfigurationInkWorkspaceAccess = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationInkWorkspaceAccess() []string {
	return []string{
		string(Windows10GeneralConfigurationInkWorkspaceAccessdisabled),
		string(Windows10GeneralConfigurationInkWorkspaceAccessenabled),
		string(Windows10GeneralConfigurationInkWorkspaceAccessnotConfigured),
	}
}

func parseWindows10GeneralConfigurationInkWorkspaceAccess(input string) (*Windows10GeneralConfigurationInkWorkspaceAccess, error) {
	vals := map[string]Windows10GeneralConfigurationInkWorkspaceAccess{
		"disabled":      Windows10GeneralConfigurationInkWorkspaceAccessdisabled,
		"enabled":       Windows10GeneralConfigurationInkWorkspaceAccessenabled,
		"notconfigured": Windows10GeneralConfigurationInkWorkspaceAccessnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationInkWorkspaceAccess(input)
	return &out, nil
}
