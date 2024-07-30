package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel string

const (
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High        AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "high"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low         AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "low"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "medium"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "notSet"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured     AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "secured"
	AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel = "unavailable"
)

func PossibleValuesForAndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel() []string {
	return []string{
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured),
		string(AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable),
	}
}

func (s *AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input string) (*AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel, error) {
	vals := map[string]AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel{
		"high":        AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_High,
		"low":         AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Low,
		"medium":      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Medium,
		"notset":      AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_NotSet,
		"secured":     AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Secured,
		"unavailable": AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyAdvancedThreatProtectionRequiredSecurityLevel(input)
	return &out, nil
}
