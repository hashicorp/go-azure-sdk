package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformCredentialAuthenticationMethodKeyStrength string

const (
	PlatformCredentialAuthenticationMethodKeyStrength_Normal  PlatformCredentialAuthenticationMethodKeyStrength = "normal"
	PlatformCredentialAuthenticationMethodKeyStrength_Unknown PlatformCredentialAuthenticationMethodKeyStrength = "unknown"
	PlatformCredentialAuthenticationMethodKeyStrength_Weak    PlatformCredentialAuthenticationMethodKeyStrength = "weak"
)

func PossibleValuesForPlatformCredentialAuthenticationMethodKeyStrength() []string {
	return []string{
		string(PlatformCredentialAuthenticationMethodKeyStrength_Normal),
		string(PlatformCredentialAuthenticationMethodKeyStrength_Unknown),
		string(PlatformCredentialAuthenticationMethodKeyStrength_Weak),
	}
}

func (s *PlatformCredentialAuthenticationMethodKeyStrength) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlatformCredentialAuthenticationMethodKeyStrength(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlatformCredentialAuthenticationMethodKeyStrength(input string) (*PlatformCredentialAuthenticationMethodKeyStrength, error) {
	vals := map[string]PlatformCredentialAuthenticationMethodKeyStrength{
		"normal":  PlatformCredentialAuthenticationMethodKeyStrength_Normal,
		"unknown": PlatformCredentialAuthenticationMethodKeyStrength_Unknown,
		"weak":    PlatformCredentialAuthenticationMethodKeyStrength_Weak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformCredentialAuthenticationMethodKeyStrength(input)
	return &out, nil
}
