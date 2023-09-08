package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh          TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "High"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevellow           TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Low"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium        TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Medium"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "NotConfigured"
	TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured       TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Secured"
)

func PossibleValuesForTargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevellow),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured),
		string(TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured),
	}
}

func parseTargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh,
		"low":           TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevellow,
		"medium":        TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium,
		"notconfigured": TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured,
		"secured":       TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}
