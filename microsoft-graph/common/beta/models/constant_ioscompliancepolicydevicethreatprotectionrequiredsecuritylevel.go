package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForIosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseIosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
