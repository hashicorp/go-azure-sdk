package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh          WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "High"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevellow           WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Low"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium        WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Medium"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "NotConfigured"
	WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured       WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Secured"
)

func PossibleValuesForWindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevellow),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured),
		string(WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured),
	}
}

func parseWindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh,
		"low":           WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevellow,
		"medium":        WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium,
		"notconfigured": WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured,
		"secured":       WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}
