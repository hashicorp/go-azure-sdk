package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls string

const (
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlschallengeHandshakeAuthenticationProtocol MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "ChallengeHandshakeAuthenticationProtocol"
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsmicrosoftChap                            MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "MicrosoftChap"
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsmicrosoftChapVersionTwo                  MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "MicrosoftChapVersionTwo"
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsunencryptedPassword                      MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "UnencryptedPassword"
)

func PossibleValuesForMacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlschallengeHandshakeAuthenticationProtocol),
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsmicrosoftChap),
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsmicrosoftChapVersionTwo),
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsunencryptedPassword),
	}
}

func parseMacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls(input string) (*MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls, error) {
	vals := map[string]MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls{
		"challengehandshakeauthenticationprotocol": MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlschallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsmicrosoftChap,
		"microsoftchapversiontwo": MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsmicrosoftChapVersionTwo,
		"unencryptedpassword":     MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtlsunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls(input)
	return &out, nil
}
