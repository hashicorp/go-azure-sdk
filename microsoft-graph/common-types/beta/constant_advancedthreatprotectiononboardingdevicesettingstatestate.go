package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingDeviceSettingStateState string

const (
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_Compliant     AdvancedThreatProtectionOnboardingDeviceSettingStateState = "compliant"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_Conflict      AdvancedThreatProtectionOnboardingDeviceSettingStateState = "conflict"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_Error         AdvancedThreatProtectionOnboardingDeviceSettingStateState = "error"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_NonCompliant  AdvancedThreatProtectionOnboardingDeviceSettingStateState = "nonCompliant"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_NotApplicable AdvancedThreatProtectionOnboardingDeviceSettingStateState = "notApplicable"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_NotAssigned   AdvancedThreatProtectionOnboardingDeviceSettingStateState = "notAssigned"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_Remediated    AdvancedThreatProtectionOnboardingDeviceSettingStateState = "remediated"
	AdvancedThreatProtectionOnboardingDeviceSettingStateState_Unknown       AdvancedThreatProtectionOnboardingDeviceSettingStateState = "unknown"
)

func PossibleValuesForAdvancedThreatProtectionOnboardingDeviceSettingStateState() []string {
	return []string{
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_Compliant),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_Conflict),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_Error),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_NonCompliant),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_NotApplicable),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_NotAssigned),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_Remediated),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateState_Unknown),
	}
}

func (s *AdvancedThreatProtectionOnboardingDeviceSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvancedThreatProtectionOnboardingDeviceSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvancedThreatProtectionOnboardingDeviceSettingStateState(input string) (*AdvancedThreatProtectionOnboardingDeviceSettingStateState, error) {
	vals := map[string]AdvancedThreatProtectionOnboardingDeviceSettingStateState{
		"compliant":     AdvancedThreatProtectionOnboardingDeviceSettingStateState_Compliant,
		"conflict":      AdvancedThreatProtectionOnboardingDeviceSettingStateState_Conflict,
		"error":         AdvancedThreatProtectionOnboardingDeviceSettingStateState_Error,
		"noncompliant":  AdvancedThreatProtectionOnboardingDeviceSettingStateState_NonCompliant,
		"notapplicable": AdvancedThreatProtectionOnboardingDeviceSettingStateState_NotApplicable,
		"notassigned":   AdvancedThreatProtectionOnboardingDeviceSettingStateState_NotAssigned,
		"remediated":    AdvancedThreatProtectionOnboardingDeviceSettingStateState_Remediated,
		"unknown":       AdvancedThreatProtectionOnboardingDeviceSettingStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedThreatProtectionOnboardingDeviceSettingStateState(input)
	return &out, nil
}
