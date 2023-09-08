package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	ManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh          ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "High"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevellow           ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Low"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium        ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Medium"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "NotConfigured"
	ManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured       ManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Secured"
)

func PossibleValuesForManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevellow),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured),
		string(ManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured),
	}
}

func parseManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*ManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]ManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          ManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh,
		"low":           ManagedAppProtectionMaximumAllowedDeviceThreatLevellow,
		"medium":        ManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium,
		"notconfigured": ManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured,
		"secured":       ManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}
