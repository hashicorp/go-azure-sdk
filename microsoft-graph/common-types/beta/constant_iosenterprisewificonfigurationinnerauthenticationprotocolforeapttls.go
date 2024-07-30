package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls string

const (
	IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "challengeHandshakeAuthenticationProtocol"
	IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap                            IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChap"
	IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo                  IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "microsoftChapVersionTwo"
	IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword                      IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls = "unencryptedPassword"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol),
		string(IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap),
		string(IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo),
		string(IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword),
	}
}

func (s *IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input string) (*IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls{
		"challengehandshakeauthenticationprotocol": IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChap,
		"microsoftchapversiontwo": IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_MicrosoftChapVersionTwo,
		"unencryptedpassword":     IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationInnerAuthenticationProtocolForEapTtls(input)
	return &out, nil
}
