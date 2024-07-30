package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInAppTokenProtectionStatus string

const (
	SignInAppTokenProtectionStatus_Bound   SignInAppTokenProtectionStatus = "bound"
	SignInAppTokenProtectionStatus_None    SignInAppTokenProtectionStatus = "none"
	SignInAppTokenProtectionStatus_Unbound SignInAppTokenProtectionStatus = "unbound"
)

func PossibleValuesForSignInAppTokenProtectionStatus() []string {
	return []string{
		string(SignInAppTokenProtectionStatus_Bound),
		string(SignInAppTokenProtectionStatus_None),
		string(SignInAppTokenProtectionStatus_Unbound),
	}
}

func (s *SignInAppTokenProtectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInAppTokenProtectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInAppTokenProtectionStatus(input string) (*SignInAppTokenProtectionStatus, error) {
	vals := map[string]SignInAppTokenProtectionStatus{
		"bound":   SignInAppTokenProtectionStatus_Bound,
		"none":    SignInAppTokenProtectionStatus_None,
		"unbound": SignInAppTokenProtectionStatus_Unbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInAppTokenProtectionStatus(input)
	return &out, nil
}
