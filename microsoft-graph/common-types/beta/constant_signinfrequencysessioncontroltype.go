package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlType string

const (
	SignInFrequencySessionControlType_Days  SignInFrequencySessionControlType = "days"
	SignInFrequencySessionControlType_Hours SignInFrequencySessionControlType = "hours"
)

func PossibleValuesForSignInFrequencySessionControlType() []string {
	return []string{
		string(SignInFrequencySessionControlType_Days),
		string(SignInFrequencySessionControlType_Hours),
	}
}

func (s *SignInFrequencySessionControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInFrequencySessionControlType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInFrequencySessionControlType(input string) (*SignInFrequencySessionControlType, error) {
	vals := map[string]SignInFrequencySessionControlType{
		"days":  SignInFrequencySessionControlType_Days,
		"hours": SignInFrequencySessionControlType_Hours,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencySessionControlType(input)
	return &out, nil
}
