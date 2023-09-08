package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication string

const (
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationnone                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "None"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationoath                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "Oath"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationpush                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "Push"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationsms                  UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "Sms"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceAlternateMobile UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "VoiceAlternateMobile"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceMobile          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "VoiceMobile"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceOffice          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "VoiceOffice"
)

func PossibleValuesForUserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication() []string {
	return []string{
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationnone),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationoath),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationpush),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationsms),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceAlternateMobile),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceMobile),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceOffice),
	}
}

func parseUserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication(input string) (*UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication, error) {
	vals := map[string]UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication{
		"none":                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationnone,
		"oath":                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationoath,
		"push":                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationpush,
		"sms":                  UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationsms,
		"voicealternatemobile": UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceAlternateMobile,
		"voicemobile":          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceMobile,
		"voiceoffice":          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthenticationvoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication(input)
	return &out, nil
}
