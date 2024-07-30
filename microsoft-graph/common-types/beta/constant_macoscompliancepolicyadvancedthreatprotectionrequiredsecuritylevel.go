package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High        MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "high"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low         MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "low"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "medium"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "notSet"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured     MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "secured"
	MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForMacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured),
		string(MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High,
		"low":         MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low,
		"medium":      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
