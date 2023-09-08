package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage string

const (
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsageallowed       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "Allowed"
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsageblocked       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "Blocked"
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsagenotConfigured WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "NotConfigured"
	WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsagerequired      WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage = "Required"
)

func PossibleValuesForWindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage() []string {
	return []string{
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsageallowed),
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsageblocked),
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsagenotConfigured),
		string(WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsagerequired),
	}
}

func parseWindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage(input string) (*WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage, error) {
	vals := map[string]WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage{
		"allowed":       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsageallowed,
		"blocked":       WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsageblocked,
		"notconfigured": WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsagenotConfigured,
		"required":      WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsagerequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage(input)
	return &out, nil
}
