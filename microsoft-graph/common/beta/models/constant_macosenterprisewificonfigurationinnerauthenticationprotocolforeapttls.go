package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "ChallengeHandshakeAuthenticationProtocol"
	MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap                            MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChap"
	MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo                  MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChapVersionTwo"
	MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword                      MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "UnencryptedPassword"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol),
		string(MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap),
		string(MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo),
		string(MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword),
	}
}

func parseMacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap,
		"microsoftchapversiontwo": MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo,
		"unencryptedpassword":     MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
