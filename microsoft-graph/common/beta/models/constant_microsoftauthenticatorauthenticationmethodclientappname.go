package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodClientAppName string

const (
	MicrosoftAuthenticatorAuthenticationMethodClientAppNamemicrosoftAuthenticator MicrosoftAuthenticatorAuthenticationMethodClientAppName = "MicrosoftAuthenticator"
	MicrosoftAuthenticatorAuthenticationMethodClientAppNameoutlookMobile          MicrosoftAuthenticatorAuthenticationMethodClientAppName = "OutlookMobile"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodClientAppName() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodClientAppNamemicrosoftAuthenticator),
		string(MicrosoftAuthenticatorAuthenticationMethodClientAppNameoutlookMobile),
	}
}

func parseMicrosoftAuthenticatorAuthenticationMethodClientAppName(input string) (*MicrosoftAuthenticatorAuthenticationMethodClientAppName, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodClientAppName{
		"microsoftauthenticator": MicrosoftAuthenticatorAuthenticationMethodClientAppNamemicrosoftAuthenticator,
		"outlookmobile":          MicrosoftAuthenticatorAuthenticationMethodClientAppNameoutlookMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodClientAppName(input)
	return &out, nil
}
