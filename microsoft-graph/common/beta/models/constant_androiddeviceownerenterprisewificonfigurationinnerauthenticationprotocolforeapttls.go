package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "ChallengeHandshakeAuthenticationProtocol"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap                            AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChap"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo                  AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChapVersionTwo"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword                      AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "UnencryptedPassword"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword),
	}
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap,
		"microsoftchapversiontwo": AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
