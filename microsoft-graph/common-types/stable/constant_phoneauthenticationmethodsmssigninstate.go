package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneAuthenticationMethodSmsSignInState string

const (
	PhoneAuthenticationMethodSmsSignInState_NotAllowedByPolicy   PhoneAuthenticationMethodSmsSignInState = "notAllowedByPolicy"
	PhoneAuthenticationMethodSmsSignInState_NotConfigured        PhoneAuthenticationMethodSmsSignInState = "notConfigured"
	PhoneAuthenticationMethodSmsSignInState_NotEnabled           PhoneAuthenticationMethodSmsSignInState = "notEnabled"
	PhoneAuthenticationMethodSmsSignInState_NotSupported         PhoneAuthenticationMethodSmsSignInState = "notSupported"
	PhoneAuthenticationMethodSmsSignInState_PhoneNumberNotUnique PhoneAuthenticationMethodSmsSignInState = "phoneNumberNotUnique"
	PhoneAuthenticationMethodSmsSignInState_Ready                PhoneAuthenticationMethodSmsSignInState = "ready"
)

func PossibleValuesForPhoneAuthenticationMethodSmsSignInState() []string {
	return []string{
		string(PhoneAuthenticationMethodSmsSignInState_NotAllowedByPolicy),
		string(PhoneAuthenticationMethodSmsSignInState_NotConfigured),
		string(PhoneAuthenticationMethodSmsSignInState_NotEnabled),
		string(PhoneAuthenticationMethodSmsSignInState_NotSupported),
		string(PhoneAuthenticationMethodSmsSignInState_PhoneNumberNotUnique),
		string(PhoneAuthenticationMethodSmsSignInState_Ready),
	}
}

func (s *PhoneAuthenticationMethodSmsSignInState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhoneAuthenticationMethodSmsSignInState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhoneAuthenticationMethodSmsSignInState(input string) (*PhoneAuthenticationMethodSmsSignInState, error) {
	vals := map[string]PhoneAuthenticationMethodSmsSignInState{
		"notallowedbypolicy":   PhoneAuthenticationMethodSmsSignInState_NotAllowedByPolicy,
		"notconfigured":        PhoneAuthenticationMethodSmsSignInState_NotConfigured,
		"notenabled":           PhoneAuthenticationMethodSmsSignInState_NotEnabled,
		"notsupported":         PhoneAuthenticationMethodSmsSignInState_NotSupported,
		"phonenumbernotunique": PhoneAuthenticationMethodSmsSignInState_PhoneNumberNotUnique,
		"ready":                PhoneAuthenticationMethodSmsSignInState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneAuthenticationMethodSmsSignInState(input)
	return &out, nil
}
