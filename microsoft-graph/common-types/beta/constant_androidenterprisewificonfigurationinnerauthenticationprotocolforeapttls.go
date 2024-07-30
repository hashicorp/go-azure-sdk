package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "challengeHandshakeAuthenticationProtocol"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap                            AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChap"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo                  AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChapVersionTwo"
	AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword                      AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "unencryptedPassword"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo),
		string(AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword),
	}
}

func (s *AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap,
		"microsoftchapversiontwo": AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
