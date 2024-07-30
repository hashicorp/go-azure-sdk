package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentInfoState string

const (
	SecurityContentInfoState_Motion SecurityContentInfoState = "motion"
	SecurityContentInfoState_Rest   SecurityContentInfoState = "rest"
	SecurityContentInfoState_Use    SecurityContentInfoState = "use"
)

func PossibleValuesForSecurityContentInfoState() []string {
	return []string{
		string(SecurityContentInfoState_Motion),
		string(SecurityContentInfoState_Rest),
		string(SecurityContentInfoState_Use),
	}
}

func (s *SecurityContentInfoState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContentInfoState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContentInfoState(input string) (*SecurityContentInfoState, error) {
	vals := map[string]SecurityContentInfoState{
		"motion": SecurityContentInfoState_Motion,
		"rest":   SecurityContentInfoState_Rest,
		"use":    SecurityContentInfoState_Use,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentInfoState(input)
	return &out, nil
}
