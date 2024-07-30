package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS string

const (
	WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_ChallengeHandshakeAuthenticationProtocol WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS = "challengeHandshakeAuthenticationProtocol"
	WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChap                            WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS = "microsoftChap"
	WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChapVersionTwo                  WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS = "microsoftChapVersionTwo"
	WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_UnencryptedPassword                      WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS = "unencryptedPassword"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_ChallengeHandshakeAuthenticationProtocol),
		string(WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChap),
		string(WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChapVersionTwo),
		string(WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_UnencryptedPassword),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS(input string) (*WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS{
		"challengehandshakeauthenticationprotocol": WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChap,
		"microsoftchapversiontwo": WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChapVersionTwo,
		"unencryptedpassword":     WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationInnerAuthenticationProtocolForEAPTTLS(input)
	return &out, nil
}
