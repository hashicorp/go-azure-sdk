package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInSignInTokenProtectionStatus string

const (
	SignInSignInTokenProtectionStatus_Bound   SignInSignInTokenProtectionStatus = "bound"
	SignInSignInTokenProtectionStatus_None    SignInSignInTokenProtectionStatus = "none"
	SignInSignInTokenProtectionStatus_Unbound SignInSignInTokenProtectionStatus = "unbound"
)

func PossibleValuesForSignInSignInTokenProtectionStatus() []string {
	return []string{
		string(SignInSignInTokenProtectionStatus_Bound),
		string(SignInSignInTokenProtectionStatus_None),
		string(SignInSignInTokenProtectionStatus_Unbound),
	}
}

func (s *SignInSignInTokenProtectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInSignInTokenProtectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInSignInTokenProtectionStatus(input string) (*SignInSignInTokenProtectionStatus, error) {
	vals := map[string]SignInSignInTokenProtectionStatus{
		"bound":   SignInSignInTokenProtectionStatus_Bound,
		"none":    SignInSignInTokenProtectionStatus_None,
		"unbound": SignInSignInTokenProtectionStatus_Unbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInSignInTokenProtectionStatus(input)
	return &out, nil
}
