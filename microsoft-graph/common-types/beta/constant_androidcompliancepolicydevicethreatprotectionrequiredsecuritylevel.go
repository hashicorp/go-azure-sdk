package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
