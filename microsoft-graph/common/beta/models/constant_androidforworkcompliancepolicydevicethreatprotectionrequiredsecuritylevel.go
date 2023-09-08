package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
