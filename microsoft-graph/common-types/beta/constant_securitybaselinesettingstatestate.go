package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineSettingStateState string

const (
	SecurityBaselineSettingStateState_Conflict      SecurityBaselineSettingStateState = "conflict"
	SecurityBaselineSettingStateState_Error         SecurityBaselineSettingStateState = "error"
	SecurityBaselineSettingStateState_NotApplicable SecurityBaselineSettingStateState = "notApplicable"
	SecurityBaselineSettingStateState_NotSecure     SecurityBaselineSettingStateState = "notSecure"
	SecurityBaselineSettingStateState_Secure        SecurityBaselineSettingStateState = "secure"
	SecurityBaselineSettingStateState_Unknown       SecurityBaselineSettingStateState = "unknown"
)

func PossibleValuesForSecurityBaselineSettingStateState() []string {
	return []string{
		string(SecurityBaselineSettingStateState_Conflict),
		string(SecurityBaselineSettingStateState_Error),
		string(SecurityBaselineSettingStateState_NotApplicable),
		string(SecurityBaselineSettingStateState_NotSecure),
		string(SecurityBaselineSettingStateState_Secure),
		string(SecurityBaselineSettingStateState_Unknown),
	}
}

func (s *SecurityBaselineSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineSettingStateState(input string) (*SecurityBaselineSettingStateState, error) {
	vals := map[string]SecurityBaselineSettingStateState{
		"conflict":      SecurityBaselineSettingStateState_Conflict,
		"error":         SecurityBaselineSettingStateState_Error,
		"notapplicable": SecurityBaselineSettingStateState_NotApplicable,
		"notsecure":     SecurityBaselineSettingStateState_NotSecure,
		"secure":        SecurityBaselineSettingStateState_Secure,
		"unknown":       SecurityBaselineSettingStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineSettingStateState(input)
	return &out, nil
}
