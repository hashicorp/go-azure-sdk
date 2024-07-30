package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "challengeHandshakeAuthenticationProtocol"
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap                            AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChap"
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo                  AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChapVersionTwo"
	AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword                      AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "unencryptedPassword"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol),
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap),
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo),
		string(AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword),
	}
}

func (s *AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap,
		"microsoftchapversiontwo": AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
