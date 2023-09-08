package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "ChallengeHandshakeAuthenticationProtocol"
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap                            AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChap"
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo                  AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "MicrosoftChapVersionTwo"
	AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword                      AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "UnencryptedPassword"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword),
	}
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlschallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChap,
		"microsoftchapversiontwo": AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsmicrosoftChapVersionTwo,
		"unencryptedpassword":     AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtlsunencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
