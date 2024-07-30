package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlFrequencyInterval string

const (
	SignInFrequencySessionControlFrequencyInterval_EveryTime SignInFrequencySessionControlFrequencyInterval = "everyTime"
	SignInFrequencySessionControlFrequencyInterval_TimeBased SignInFrequencySessionControlFrequencyInterval = "timeBased"
)

func PossibleValuesForSignInFrequencySessionControlFrequencyInterval() []string {
	return []string{
		string(SignInFrequencySessionControlFrequencyInterval_EveryTime),
		string(SignInFrequencySessionControlFrequencyInterval_TimeBased),
	}
}

func (s *SignInFrequencySessionControlFrequencyInterval) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInFrequencySessionControlFrequencyInterval(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInFrequencySessionControlFrequencyInterval(input string) (*SignInFrequencySessionControlFrequencyInterval, error) {
	vals := map[string]SignInFrequencySessionControlFrequencyInterval{
		"everytime": SignInFrequencySessionControlFrequencyInterval_EveryTime,
		"timebased": SignInFrequencySessionControlFrequencyInterval_TimeBased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencySessionControlFrequencyInterval(input)
	return &out, nil
}
