package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
