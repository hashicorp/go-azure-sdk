package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneAuthenticationMethodSmsSignInState string

const (
	PhoneAuthenticationMethodSmsSignInStatenotAllowedByPolicy   PhoneAuthenticationMethodSmsSignInState = "NotAllowedByPolicy"
	PhoneAuthenticationMethodSmsSignInStatenotConfigured        PhoneAuthenticationMethodSmsSignInState = "NotConfigured"
	PhoneAuthenticationMethodSmsSignInStatenotEnabled           PhoneAuthenticationMethodSmsSignInState = "NotEnabled"
	PhoneAuthenticationMethodSmsSignInStatenotSupported         PhoneAuthenticationMethodSmsSignInState = "NotSupported"
	PhoneAuthenticationMethodSmsSignInStatephoneNumberNotUnique PhoneAuthenticationMethodSmsSignInState = "PhoneNumberNotUnique"
	PhoneAuthenticationMethodSmsSignInStateready                PhoneAuthenticationMethodSmsSignInState = "Ready"
)

func PossibleValuesForPhoneAuthenticationMethodSmsSignInState() []string {
	return []string{
		string(PhoneAuthenticationMethodSmsSignInStatenotAllowedByPolicy),
		string(PhoneAuthenticationMethodSmsSignInStatenotConfigured),
		string(PhoneAuthenticationMethodSmsSignInStatenotEnabled),
		string(PhoneAuthenticationMethodSmsSignInStatenotSupported),
		string(PhoneAuthenticationMethodSmsSignInStatephoneNumberNotUnique),
		string(PhoneAuthenticationMethodSmsSignInStateready),
	}
}

func parsePhoneAuthenticationMethodSmsSignInState(input string) (*PhoneAuthenticationMethodSmsSignInState, error) {
	vals := map[string]PhoneAuthenticationMethodSmsSignInState{
		"notallowedbypolicy":   PhoneAuthenticationMethodSmsSignInStatenotAllowedByPolicy,
		"notconfigured":        PhoneAuthenticationMethodSmsSignInStatenotConfigured,
		"notenabled":           PhoneAuthenticationMethodSmsSignInStatenotEnabled,
		"notsupported":         PhoneAuthenticationMethodSmsSignInStatenotSupported,
		"phonenumbernotunique": PhoneAuthenticationMethodSmsSignInStatephoneNumberNotUnique,
		"ready":                PhoneAuthenticationMethodSmsSignInStateready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneAuthenticationMethodSmsSignInState(input)
	return &out, nil
}
