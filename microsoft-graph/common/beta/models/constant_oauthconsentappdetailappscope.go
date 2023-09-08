package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuthConsentAppDetailAppScope string

const (
	OAuthConsentAppDetailAppScopereadAllChat      OAuthConsentAppDetailAppScope = "ReadAllChat"
	OAuthConsentAppDetailAppScopereadAllFile      OAuthConsentAppDetailAppScope = "ReadAllFile"
	OAuthConsentAppDetailAppScopereadAndWriteMail OAuthConsentAppDetailAppScope = "ReadAndWriteMail"
	OAuthConsentAppDetailAppScopereadCalendar     OAuthConsentAppDetailAppScope = "ReadCalendar"
	OAuthConsentAppDetailAppScopereadContact      OAuthConsentAppDetailAppScope = "ReadContact"
	OAuthConsentAppDetailAppScopereadMail         OAuthConsentAppDetailAppScope = "ReadMail"
	OAuthConsentAppDetailAppScopesendMail         OAuthConsentAppDetailAppScope = "SendMail"
	OAuthConsentAppDetailAppScopeunknown          OAuthConsentAppDetailAppScope = "Unknown"
)

func PossibleValuesForOAuthConsentAppDetailAppScope() []string {
	return []string{
		string(OAuthConsentAppDetailAppScopereadAllChat),
		string(OAuthConsentAppDetailAppScopereadAllFile),
		string(OAuthConsentAppDetailAppScopereadAndWriteMail),
		string(OAuthConsentAppDetailAppScopereadCalendar),
		string(OAuthConsentAppDetailAppScopereadContact),
		string(OAuthConsentAppDetailAppScopereadMail),
		string(OAuthConsentAppDetailAppScopesendMail),
		string(OAuthConsentAppDetailAppScopeunknown),
	}
}

func parseOAuthConsentAppDetailAppScope(input string) (*OAuthConsentAppDetailAppScope, error) {
	vals := map[string]OAuthConsentAppDetailAppScope{
		"readallchat":      OAuthConsentAppDetailAppScopereadAllChat,
		"readallfile":      OAuthConsentAppDetailAppScopereadAllFile,
		"readandwritemail": OAuthConsentAppDetailAppScopereadAndWriteMail,
		"readcalendar":     OAuthConsentAppDetailAppScopereadCalendar,
		"readcontact":      OAuthConsentAppDetailAppScopereadContact,
		"readmail":         OAuthConsentAppDetailAppScopereadMail,
		"sendmail":         OAuthConsentAppDetailAppScopesendMail,
		"unknown":          OAuthConsentAppDetailAppScopeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OAuthConsentAppDetailAppScope(input)
	return &out, nil
}
