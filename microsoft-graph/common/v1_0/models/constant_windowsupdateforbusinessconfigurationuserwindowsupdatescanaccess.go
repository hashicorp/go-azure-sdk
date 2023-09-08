package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess string

const (
	WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessdisabled      WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess = "Disabled"
	WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessenabled       WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess = "Enabled"
	WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessnotConfigured WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess = "NotConfigured"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessdisabled),
		string(WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessenabled),
		string(WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessnotConfigured),
	}
}

func parseWindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess(input string) (*WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess{
		"disabled":      WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessdisabled,
		"enabled":       WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessenabled,
		"notconfigured": WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccessnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess(input)
	return &out, nil
}
