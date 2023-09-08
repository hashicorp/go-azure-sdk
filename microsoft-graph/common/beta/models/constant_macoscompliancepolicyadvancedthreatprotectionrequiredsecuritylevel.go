package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh        MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "High"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow         MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Low"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Medium"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "NotSet"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured     MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Secured"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForMacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseMacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh,
		"low":         MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow,
		"medium":      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium,
		"notset":      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
