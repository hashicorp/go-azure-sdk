package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh        IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "High"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow         IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Low"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Medium"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "NotSet"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured     IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Secured"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForIosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseIosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh,
		"low":         IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow,
		"medium":      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium,
		"notset":      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
