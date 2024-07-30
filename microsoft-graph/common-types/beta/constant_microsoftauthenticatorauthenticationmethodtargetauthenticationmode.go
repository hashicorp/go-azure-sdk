package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode string

const (
	MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_Any             MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode = "any"
	MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_DeviceBasedPush MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode = "deviceBasedPush"
	MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_Push            MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode = "push"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_Any),
		string(MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_DeviceBasedPush),
		string(MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_Push),
	}
}

func (s *MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode(input string) (*MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode{
		"any":             MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_Any,
		"devicebasedpush": MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_DeviceBasedPush,
		"push":            MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode_Push,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodTargetAuthenticationMode(input)
	return &out, nil
}
