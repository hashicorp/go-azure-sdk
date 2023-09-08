package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh        AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "High"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow         AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Low"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Medium"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "NotSet"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured     AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Secured"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "Unavailable"
)

func PossibleValuesForAndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable),
	}
}

func parseAndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelhigh,
		"low":         AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevellow,
		"medium":      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelmedium,
		"notset":      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelnotSet,
		"secured":     AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelsecured,
		"unavailable": AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevelunavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
