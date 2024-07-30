package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuthConsentAppDetailAppScope string

const (
	OAuthConsentAppDetailAppScope_ReadAllChat      OAuthConsentAppDetailAppScope = "readAllChat"
	OAuthConsentAppDetailAppScope_ReadAllFile      OAuthConsentAppDetailAppScope = "readAllFile"
	OAuthConsentAppDetailAppScope_ReadAndWriteMail OAuthConsentAppDetailAppScope = "readAndWriteMail"
	OAuthConsentAppDetailAppScope_ReadCalendar     OAuthConsentAppDetailAppScope = "readCalendar"
	OAuthConsentAppDetailAppScope_ReadContact      OAuthConsentAppDetailAppScope = "readContact"
	OAuthConsentAppDetailAppScope_ReadMail         OAuthConsentAppDetailAppScope = "readMail"
	OAuthConsentAppDetailAppScope_SendMail         OAuthConsentAppDetailAppScope = "sendMail"
	OAuthConsentAppDetailAppScope_Unknown          OAuthConsentAppDetailAppScope = "unknown"
)

func PossibleValuesForOAuthConsentAppDetailAppScope() []string {
	return []string{
		string(OAuthConsentAppDetailAppScope_ReadAllChat),
		string(OAuthConsentAppDetailAppScope_ReadAllFile),
		string(OAuthConsentAppDetailAppScope_ReadAndWriteMail),
		string(OAuthConsentAppDetailAppScope_ReadCalendar),
		string(OAuthConsentAppDetailAppScope_ReadContact),
		string(OAuthConsentAppDetailAppScope_ReadMail),
		string(OAuthConsentAppDetailAppScope_SendMail),
		string(OAuthConsentAppDetailAppScope_Unknown),
	}
}

func (s *OAuthConsentAppDetailAppScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOAuthConsentAppDetailAppScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOAuthConsentAppDetailAppScope(input string) (*OAuthConsentAppDetailAppScope, error) {
	vals := map[string]OAuthConsentAppDetailAppScope{
		"readallchat":      OAuthConsentAppDetailAppScope_ReadAllChat,
		"readallfile":      OAuthConsentAppDetailAppScope_ReadAllFile,
		"readandwritemail": OAuthConsentAppDetailAppScope_ReadAndWriteMail,
		"readcalendar":     OAuthConsentAppDetailAppScope_ReadCalendar,
		"readcontact":      OAuthConsentAppDetailAppScope_ReadContact,
		"readmail":         OAuthConsentAppDetailAppScope_ReadMail,
		"sendmail":         OAuthConsentAppDetailAppScope_SendMail,
		"unknown":          OAuthConsentAppDetailAppScope_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OAuthConsentAppDetailAppScope(input)
	return &out, nil
}
