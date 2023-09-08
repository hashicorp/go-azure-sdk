package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh        AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "High"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow         AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Low"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured     AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
