package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
