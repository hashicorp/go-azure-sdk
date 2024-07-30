package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "challengeHandshakeAuthenticationProtocol"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap                            AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChap"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo                  AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChapVersionTwo"
	AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword                      AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "unencryptedPassword"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword),
	}
}

func (s *AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap,
		"microsoftchapversiontwo": AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo,
		"unencryptedpassword":     AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
