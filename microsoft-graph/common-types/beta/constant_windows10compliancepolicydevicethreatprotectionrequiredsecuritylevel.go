package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForWindows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
