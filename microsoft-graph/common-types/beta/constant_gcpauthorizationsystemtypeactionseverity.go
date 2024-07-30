package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpAuthorizationSystemTypeActionSeverity string

const (
	GcpAuthorizationSystemTypeActionSeverity_High   GcpAuthorizationSystemTypeActionSeverity = "high"
	GcpAuthorizationSystemTypeActionSeverity_Normal GcpAuthorizationSystemTypeActionSeverity = "normal"
)

func PossibleValuesForGcpAuthorizationSystemTypeActionSeverity() []string {
	return []string{
		string(GcpAuthorizationSystemTypeActionSeverity_High),
		string(GcpAuthorizationSystemTypeActionSeverity_Normal),
	}
}

func (s *GcpAuthorizationSystemTypeActionSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGcpAuthorizationSystemTypeActionSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGcpAuthorizationSystemTypeActionSeverity(input string) (*GcpAuthorizationSystemTypeActionSeverity, error) {
	vals := map[string]GcpAuthorizationSystemTypeActionSeverity{
		"high":   GcpAuthorizationSystemTypeActionSeverity_High,
		"normal": GcpAuthorizationSystemTypeActionSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GcpAuthorizationSystemTypeActionSeverity(input)
	return &out, nil
}
