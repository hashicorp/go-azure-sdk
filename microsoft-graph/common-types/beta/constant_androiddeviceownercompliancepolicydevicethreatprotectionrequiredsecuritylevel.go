package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
