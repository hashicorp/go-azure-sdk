package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh        AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "High"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow         AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Low"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured     AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
