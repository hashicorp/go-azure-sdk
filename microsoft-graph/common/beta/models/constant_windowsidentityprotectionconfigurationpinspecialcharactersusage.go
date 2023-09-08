package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage string

const (
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsageallowed       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "Allowed"
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsageblocked       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "Blocked"
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsagenotConfigured WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "NotConfigured"
	WindowsIdentityProtectionConfigurationPinSpecialCharactersUsagerequired      WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage = "Required"
)

func PossibleValuesForWindowsIdentityProtectionConfigurationPinSpecialCharactersUsage() []string {
	return []string{
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsageallowed),
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsageblocked),
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsagenotConfigured),
		string(WindowsIdentityProtectionConfigurationPinSpecialCharactersUsagerequired),
	}
}

func parseWindowsIdentityProtectionConfigurationPinSpecialCharactersUsage(input string) (*WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage, error) {
	vals := map[string]WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage{
		"allowed":       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsageallowed,
		"blocked":       WindowsIdentityProtectionConfigurationPinSpecialCharactersUsageblocked,
		"notconfigured": WindowsIdentityProtectionConfigurationPinSpecialCharactersUsagenotConfigured,
		"required":      WindowsIdentityProtectionConfigurationPinSpecialCharactersUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage(input)
	return &out, nil
}
