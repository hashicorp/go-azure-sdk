package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForWindows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseWindows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
