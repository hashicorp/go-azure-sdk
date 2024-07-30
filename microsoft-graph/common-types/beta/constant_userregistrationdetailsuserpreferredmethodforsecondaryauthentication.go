package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication string

const (
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_None                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "none"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Oath                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "oath"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Push                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "push"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Sms                  UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "sms"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceAlternateMobile UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "voiceAlternateMobile"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceMobile          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "voiceMobile"
	UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceOffice          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication = "voiceOffice"
)

func PossibleValuesForUserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication() []string {
	return []string{
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_None),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Oath),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Push),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Sms),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceAlternateMobile),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceMobile),
		string(UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceOffice),
	}
}

func (s *UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication(input string) (*UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication, error) {
	vals := map[string]UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication{
		"none":                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_None,
		"oath":                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Oath,
		"push":                 UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Push,
		"sms":                  UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_Sms,
		"voicealternatemobile": UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceAlternateMobile,
		"voicemobile":          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceMobile,
		"voiceoffice":          UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication_VoiceOffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication(input)
	return &out, nil
}
