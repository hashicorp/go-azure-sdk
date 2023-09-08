package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForMacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseMacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
