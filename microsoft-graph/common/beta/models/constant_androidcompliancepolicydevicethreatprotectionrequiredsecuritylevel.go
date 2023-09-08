package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
