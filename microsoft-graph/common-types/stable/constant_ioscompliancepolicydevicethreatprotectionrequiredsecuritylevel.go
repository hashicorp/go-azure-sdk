package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel string

const (
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High        IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "high"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low         IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "low"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "medium"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "notSet"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured     IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "secured"
	IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForIosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured),
		string(IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input string) (*IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel{
		"high":        IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_High,
		"low":         IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Low,
		"medium":      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCompliancePolicyDeviceThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
