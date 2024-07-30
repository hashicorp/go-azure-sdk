package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "challengeHandshakeAuthenticationProtocol"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap                            AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChap"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo                  AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChapVersionTwo"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword                      AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "unencryptedPassword"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword),
	}
}

func (s *AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap,
		"microsoftchapversiontwo": AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
