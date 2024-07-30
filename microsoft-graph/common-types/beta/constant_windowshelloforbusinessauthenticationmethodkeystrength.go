package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHelloForBusinessAuthenticationMethodKeyStrength string

const (
	WindowsHelloForBusinessAuthenticationMethodKeyStrength_Normal  WindowsHelloForBusinessAuthenticationMethodKeyStrength = "normal"
	WindowsHelloForBusinessAuthenticationMethodKeyStrength_Unknown WindowsHelloForBusinessAuthenticationMethodKeyStrength = "unknown"
	WindowsHelloForBusinessAuthenticationMethodKeyStrength_Weak    WindowsHelloForBusinessAuthenticationMethodKeyStrength = "weak"
)

func PossibleValuesForWindowsHelloForBusinessAuthenticationMethodKeyStrength() []string {
	return []string{
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrength_Normal),
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrength_Unknown),
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrength_Weak),
	}
}

func (s *WindowsHelloForBusinessAuthenticationMethodKeyStrength) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsHelloForBusinessAuthenticationMethodKeyStrength(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsHelloForBusinessAuthenticationMethodKeyStrength(input string) (*WindowsHelloForBusinessAuthenticationMethodKeyStrength, error) {
	vals := map[string]WindowsHelloForBusinessAuthenticationMethodKeyStrength{
		"normal":  WindowsHelloForBusinessAuthenticationMethodKeyStrength_Normal,
		"unknown": WindowsHelloForBusinessAuthenticationMethodKeyStrength_Unknown,
		"weak":    WindowsHelloForBusinessAuthenticationMethodKeyStrength_Weak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHelloForBusinessAuthenticationMethodKeyStrength(input)
	return &out, nil
}
