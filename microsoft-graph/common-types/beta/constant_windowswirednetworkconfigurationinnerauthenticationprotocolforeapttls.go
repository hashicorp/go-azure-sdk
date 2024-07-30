package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS string

const (
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_ChallengeHandshakeAuthenticationProtocol WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "challengeHandshakeAuthenticationProtocol"
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChap                            WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "microsoftChap"
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChapVersionTwo                  WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "microsoftChapVersionTwo"
	WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_UnencryptedPassword                      WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS = "unencryptedPassword"
)

func PossibleValuesForWindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_ChallengeHandshakeAuthenticationProtocol),
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChap),
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChapVersionTwo),
		string(WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_UnencryptedPassword),
	}
}

func (s *WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS(input string) (*WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS, error) {
	vals := map[string]WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS{
		"challengehandshakeauthenticationprotocol": WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChap,
		"microsoftchapversiontwo": WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_MicrosoftChapVersionTwo,
		"unencryptedpassword":     WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationInnerAuthenticationProtocolForEAPTTLS(input)
	return &out, nil
}
