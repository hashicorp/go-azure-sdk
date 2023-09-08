package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionMaximumAllowedDeviceThreatLevel string

const (
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh          IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "High"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevellow           IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Low"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium        IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Medium"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "NotConfigured"
	IosManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured       IosManagedAppProtectionMaximumAllowedDeviceThreatLevel = "Secured"
)

func PossibleValuesForIosManagedAppProtectionMaximumAllowedDeviceThreatLevel() []string {
	return []string{
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevellow),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured),
		string(IosManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured),
	}
}

func parseIosManagedAppProtectionMaximumAllowedDeviceThreatLevel(input string) (*IosManagedAppProtectionMaximumAllowedDeviceThreatLevel, error) {
	vals := map[string]IosManagedAppProtectionMaximumAllowedDeviceThreatLevel{
		"high":          IosManagedAppProtectionMaximumAllowedDeviceThreatLevelhigh,
		"low":           IosManagedAppProtectionMaximumAllowedDeviceThreatLevellow,
		"medium":        IosManagedAppProtectionMaximumAllowedDeviceThreatLevelmedium,
		"notconfigured": IosManagedAppProtectionMaximumAllowedDeviceThreatLevelnotConfigured,
		"secured":       IosManagedAppProtectionMaximumAllowedDeviceThreatLevelsecured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionMaximumAllowedDeviceThreatLevel(input)
	return &out, nil
}
