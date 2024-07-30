package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsAuthorizationSystemTypeActionSeverity string

const (
	AwsAuthorizationSystemTypeActionSeverity_High   AwsAuthorizationSystemTypeActionSeverity = "high"
	AwsAuthorizationSystemTypeActionSeverity_Normal AwsAuthorizationSystemTypeActionSeverity = "normal"
)

func PossibleValuesForAwsAuthorizationSystemTypeActionSeverity() []string {
	return []string{
		string(AwsAuthorizationSystemTypeActionSeverity_High),
		string(AwsAuthorizationSystemTypeActionSeverity_Normal),
	}
}

func (s *AwsAuthorizationSystemTypeActionSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsAuthorizationSystemTypeActionSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsAuthorizationSystemTypeActionSeverity(input string) (*AwsAuthorizationSystemTypeActionSeverity, error) {
	vals := map[string]AwsAuthorizationSystemTypeActionSeverity{
		"high":   AwsAuthorizationSystemTypeActionSeverity_High,
		"normal": AwsAuthorizationSystemTypeActionSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsAuthorizationSystemTypeActionSeverity(input)
	return &out, nil
}
