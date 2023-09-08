package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUserPauseAccess string

const (
	WindowsUpdateForBusinessConfigurationUserPauseAccessdisabled      WindowsUpdateForBusinessConfigurationUserPauseAccess = "Disabled"
	WindowsUpdateForBusinessConfigurationUserPauseAccessenabled       WindowsUpdateForBusinessConfigurationUserPauseAccess = "Enabled"
	WindowsUpdateForBusinessConfigurationUserPauseAccessnotConfigured WindowsUpdateForBusinessConfigurationUserPauseAccess = "NotConfigured"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUserPauseAccess() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUserPauseAccessdisabled),
		string(WindowsUpdateForBusinessConfigurationUserPauseAccessenabled),
		string(WindowsUpdateForBusinessConfigurationUserPauseAccessnotConfigured),
	}
}

func parseWindowsUpdateForBusinessConfigurationUserPauseAccess(input string) (*WindowsUpdateForBusinessConfigurationUserPauseAccess, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUserPauseAccess{
		"disabled":      WindowsUpdateForBusinessConfigurationUserPauseAccessdisabled,
		"enabled":       WindowsUpdateForBusinessConfigurationUserPauseAccessenabled,
		"notconfigured": WindowsUpdateForBusinessConfigurationUserPauseAccessnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUserPauseAccess(input)
	return &out, nil
}
