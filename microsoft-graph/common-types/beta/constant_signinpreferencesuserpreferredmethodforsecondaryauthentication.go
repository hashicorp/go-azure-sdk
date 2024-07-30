package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInPreferencesUserPreferredMethodForSecondaryAuthentication string

const (
	SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Oath                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "oath"
	SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Push                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "push"
	SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Sms                  SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "sms"
	SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceAlternateMobile SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "voiceAlternateMobile"
	SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceMobile          SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "voiceMobile"
	SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceOffice          SignInPreferencesUserPreferredMethodForSecondaryAuthentication = "voiceOffice"
)

func PossibleValuesForSignInPreferencesUserPreferredMethodForSecondaryAuthentication() []string {
	return []string{
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Oath),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Push),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Sms),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceAlternateMobile),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceMobile),
		string(SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceOffice),
	}
}

func (s *SignInPreferencesUserPreferredMethodForSecondaryAuthentication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInPreferencesUserPreferredMethodForSecondaryAuthentication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInPreferencesUserPreferredMethodForSecondaryAuthentication(input string) (*SignInPreferencesUserPreferredMethodForSecondaryAuthentication, error) {
	vals := map[string]SignInPreferencesUserPreferredMethodForSecondaryAuthentication{
		"oath":                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Oath,
		"push":                 SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Push,
		"sms":                  SignInPreferencesUserPreferredMethodForSecondaryAuthentication_Sms,
		"voicealternatemobile": SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceAlternateMobile,
		"voicemobile":          SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceMobile,
		"voiceoffice":          SignInPreferencesUserPreferredMethodForSecondaryAuthentication_VoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInPreferencesUserPreferredMethodForSecondaryAuthentication(input)
	return &out, nil
}
