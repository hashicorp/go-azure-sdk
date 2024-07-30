package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineDeviceStateState string

const (
	SecurityBaselineDeviceStateState_Conflict      SecurityBaselineDeviceStateState = "conflict"
	SecurityBaselineDeviceStateState_Error         SecurityBaselineDeviceStateState = "error"
	SecurityBaselineDeviceStateState_NotApplicable SecurityBaselineDeviceStateState = "notApplicable"
	SecurityBaselineDeviceStateState_NotSecure     SecurityBaselineDeviceStateState = "notSecure"
	SecurityBaselineDeviceStateState_Secure        SecurityBaselineDeviceStateState = "secure"
	SecurityBaselineDeviceStateState_Unknown       SecurityBaselineDeviceStateState = "unknown"
)

func PossibleValuesForSecurityBaselineDeviceStateState() []string {
	return []string{
		string(SecurityBaselineDeviceStateState_Conflict),
		string(SecurityBaselineDeviceStateState_Error),
		string(SecurityBaselineDeviceStateState_NotApplicable),
		string(SecurityBaselineDeviceStateState_NotSecure),
		string(SecurityBaselineDeviceStateState_Secure),
		string(SecurityBaselineDeviceStateState_Unknown),
	}
}

func (s *SecurityBaselineDeviceStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineDeviceStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineDeviceStateState(input string) (*SecurityBaselineDeviceStateState, error) {
	vals := map[string]SecurityBaselineDeviceStateState{
		"conflict":      SecurityBaselineDeviceStateState_Conflict,
		"error":         SecurityBaselineDeviceStateState_Error,
		"notapplicable": SecurityBaselineDeviceStateState_NotApplicable,
		"notsecure":     SecurityBaselineDeviceStateState_NotSecure,
		"secure":        SecurityBaselineDeviceStateState_Secure,
		"unknown":       SecurityBaselineDeviceStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineDeviceStateState(input)
	return &out, nil
}
