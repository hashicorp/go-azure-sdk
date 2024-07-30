package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
