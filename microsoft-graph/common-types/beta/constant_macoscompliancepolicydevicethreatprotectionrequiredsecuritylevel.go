package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForMacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
