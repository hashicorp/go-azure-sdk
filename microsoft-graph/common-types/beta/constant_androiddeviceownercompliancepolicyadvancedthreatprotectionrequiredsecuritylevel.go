package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High        AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "high"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low         AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "low"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "medium"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured     AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "secured"
	AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
