package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "ChallengeHandshakeAuthenticationProtocol"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap                            AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChap"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo                  AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChapVersionTwo"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword                      AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "UnencryptedPassword"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword),
	}
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap,
		"microsoftchapversiontwo": AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
