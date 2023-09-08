package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoginPageSource string

const (
	LoginPageSourceglobal  LoginPageSource = "Global"
	LoginPageSourcetenant  LoginPageSource = "Tenant"
	LoginPageSourceunknown LoginPageSource = "Unknown"
)

func PossibleValuesForLoginPageSource() []string {
	return []string{
		string(LoginPageSourceglobal),
		string(LoginPageSourcetenant),
		string(LoginPageSourceunknown),
	}
}

func parseLoginPageSource(input string) (*LoginPageSource, error) {
	vals := map[string]LoginPageSource{
		"global":  LoginPageSourceglobal,
		"tenant":  LoginPageSourcetenant,
		"unknown": LoginPageSourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginPageSource(input)
	return &out, nil
}
