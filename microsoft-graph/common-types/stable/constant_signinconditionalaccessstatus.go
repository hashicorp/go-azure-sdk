package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInConditionalAccessStatus string

const (
	SignInConditionalAccessStatus_Failure    SignInConditionalAccessStatus = "failure"
	SignInConditionalAccessStatus_NotApplied SignInConditionalAccessStatus = "notApplied"
	SignInConditionalAccessStatus_Success    SignInConditionalAccessStatus = "success"
)

func PossibleValuesForSignInConditionalAccessStatus() []string {
	return []string{
		string(SignInConditionalAccessStatus_Failure),
		string(SignInConditionalAccessStatus_NotApplied),
		string(SignInConditionalAccessStatus_Success),
	}
}

func (s *SignInConditionalAccessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInConditionalAccessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInConditionalAccessStatus(input string) (*SignInConditionalAccessStatus, error) {
	vals := map[string]SignInConditionalAccessStatus{
		"failure":    SignInConditionalAccessStatus_Failure,
		"notapplied": SignInConditionalAccessStatus_NotApplied,
		"success":    SignInConditionalAccessStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInConditionalAccessStatus(input)
	return &out, nil
}
