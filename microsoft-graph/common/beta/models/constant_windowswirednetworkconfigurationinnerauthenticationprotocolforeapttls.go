package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS string

const (
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSchallengeHandshakeAuthenticationProtocol WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "ChallengeHandshakeAuthenticationProtocol"
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSmicrosoftChap                            WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "MicrosoftChap"
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSmicrosoftChapVersionTwo                  WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "MicrosoftChapVersionTwo"
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSunencryptedPassword                      WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "UnencryptedPassword"
)

func PossibleValuesForWindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSchallengeHandshakeAuthenticationProtocol),
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSmicrosoftChap),
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSmicrosoftChapVersionTwo),
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSunencryptedPassword),
	}
}

func parseWindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS(input string) (*WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS, error) {
	vals := map[string]WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS{
		"challengehandshakeauthenticationprotocol": WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSchallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSmicrosoftChap,
		"microsoftchapversiontwo": WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSmicrosoftChapVersionTwo,
		"unencryptedpassword":     WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLSunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS(input)
	return &out, nil
}
