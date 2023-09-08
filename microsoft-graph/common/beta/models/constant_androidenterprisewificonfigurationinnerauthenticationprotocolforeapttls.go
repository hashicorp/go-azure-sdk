package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "ChallengeHandshakeAuthenticationProtocol"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap                            AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChap"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo                  AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChapVersionTwo"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword                      AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "UnencryptedPassword"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword),
	}
}

func parseAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap,
		"microsoftchapversiontwo": AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
