package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineStateState string

const (
	SecurityBaselineStateState_Conflict      SecurityBaselineStateState = "conflict"
	SecurityBaselineStateState_Error         SecurityBaselineStateState = "error"
	SecurityBaselineStateState_NotApplicable SecurityBaselineStateState = "notApplicable"
	SecurityBaselineStateState_NotSecure     SecurityBaselineStateState = "notSecure"
	SecurityBaselineStateState_Secure        SecurityBaselineStateState = "secure"
	SecurityBaselineStateState_Unknown       SecurityBaselineStateState = "unknown"
)

func PossibleValuesForSecurityBaselineStateState() []string {
	return []string{
		string(SecurityBaselineStateState_Conflict),
		string(SecurityBaselineStateState_Error),
		string(SecurityBaselineStateState_NotApplicable),
		string(SecurityBaselineStateState_NotSecure),
		string(SecurityBaselineStateState_Secure),
		string(SecurityBaselineStateState_Unknown),
	}
}

func (s *SecurityBaselineStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineStateState(input string) (*SecurityBaselineStateState, error) {
	vals := map[string]SecurityBaselineStateState{
		"conflict":      SecurityBaselineStateState_Conflict,
		"error":         SecurityBaselineStateState_Error,
		"notapplicable": SecurityBaselineStateState_NotApplicable,
		"notsecure":     SecurityBaselineStateState_NotSecure,
		"secure":        SecurityBaselineStateState_Secure,
		"unknown":       SecurityBaselineStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineStateState(input)
	return &out, nil
}
