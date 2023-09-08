package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh        AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "High"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow         AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Low"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured     AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
