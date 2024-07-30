package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High        IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "high"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low         IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "low"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "medium"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "notSet"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured     IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "secured"
	IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForIosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured),
		string(IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High,
		"low":         IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low,
		"medium":      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
