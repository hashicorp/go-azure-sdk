package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High        AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "high"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low         AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "low"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "medium"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured     AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "secured"
	AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
