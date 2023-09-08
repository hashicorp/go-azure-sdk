package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInPreferencesUserPreferredMethodForSecondaryAuthentication string

const (
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationoath                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "Oath"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationpush                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "Push"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationsms                  SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "Sms"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceAlternateMobile SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "VoiceAlternateMobile"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceMobile          SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "VoiceMobile"
	SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceOffice          SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "VoiceOffice"
)

func PossibleValuesForSignInPreferencesUserPreferredMethodForSecondaryAuthentication() []string {
	return []string{
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationoath),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationpush),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationsms),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceAlternateMobile),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceMobile),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceOffice),
	}
}

func parseSignInPreferencesUserPreferredMethodForSecondaryAuthentication(input string) (*SignInPreferencesUserPreferredMethodForSecondaryAuthentication, error) {
	vals := map[string]SignInPreferencesUserPreferredMethodForSecondaryAuthentication{
		"oath":                 SignInPreferencesUserPreferredMethodForSecondaryAuthenticationoath,
		"push":                 SignInPreferencesUserPreferredMethodForSecondaryAuthenticationpush,
		"sms":                  SignInPreferencesUserPreferredMethodForSecondaryAuthenticationsms,
		"voicealternatemobile": SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceAlternateMobile,
		"voicemobile":          SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceMobile,
		"voiceoffice":          SignInPreferencesUserPreferredMethodForSecondaryAuthenticationvoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInPreferencesUserPreferredMethodForSecondaryAuthentication(input)
	return &out, nil
}
