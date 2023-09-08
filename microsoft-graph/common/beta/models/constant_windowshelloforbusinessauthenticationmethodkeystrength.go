package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHelloForBusinessAuthenticationMethodKeyStrength string

const (
	WindowsHelloForBusinessAuthenticationMethodKeyStrengthnormal  WindowsHelloForBusinessAuthenticationMethodKeyStrength = "Normal"
	WindowsHelloForBusinessAuthenticationMethodKeyStrengthunknown WindowsHelloForBusinessAuthenticationMethodKeyStrength = "Unknown"
	WindowsHelloForBusinessAuthenticationMethodKeyStrengthweak    WindowsHelloForBusinessAuthenticationMethodKeyStrength = "Weak"
)

func PossibleValuesForWindowsHelloForBusinessAuthenticationMethodKeyStrength() []string {
	return []string{
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrengthnormal),
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrengthunknown),
		string(WindowsHelloForBusinessAuthenticationMethodKeyStrengthweak),
	}
}

func parseWindowsHelloForBusinessAuthenticationMethodKeyStrength(input string) (*WindowsHelloForBusinessAuthenticationMethodKeyStrength, error) {
	vals := map[string]WindowsHelloForBusinessAuthenticationMethodKeyStrength{
		"normal":  WindowsHelloForBusinessAuthenticationMethodKeyStrengthnormal,
		"unknown": WindowsHelloForBusinessAuthenticationMethodKeyStrengthunknown,
		"weak":    WindowsHelloForBusinessAuthenticationMethodKeyStrengthweak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHelloForBusinessAuthenticationMethodKeyStrength(input)
	return &out, nil
}
