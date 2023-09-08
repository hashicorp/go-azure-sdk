package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh          DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "High"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevellow           DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Low"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium        DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Medium"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "NotConfigured"
	DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured       DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Secured"
)

func PossibleValuesForDefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevellow),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured),
		string(DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured),
	}
}

func parseDefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh,
		"low":           DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevellow,
		"medium":        DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium,
		"notconfigured": DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured,
		"secured":       DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}
