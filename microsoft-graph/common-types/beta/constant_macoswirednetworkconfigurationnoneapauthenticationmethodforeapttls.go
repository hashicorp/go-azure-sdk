package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls string

const (
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_ChallengeHandshakeAuthenticationProtocol MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "challengeHandshakeAuthenticationProtocol"
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_MicrosoftChap                            MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "microsoftChap"
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_MicrosoftChapVersionTwo                  MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "microsoftChapVersionTwo"
	MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_UnencryptedPassword                      MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls = "unencryptedPassword"
)

func PossibleValuesForMacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_ChallengeHandshakeAuthenticationProtocol),
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_MicrosoftChap),
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_MicrosoftChapVersionTwo),
		string(MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_UnencryptedPassword),
	}
}

func (s *MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls(input string) (*MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls, error) {
	vals := map[string]MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls{
		"challengehandshakeauthenticationprotocol": MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_ChallengeHandshakeAuthenticationProtocol,
		"microsoftchap":           MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_MicrosoftChap,
		"microsoftchapversiontwo": MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_MicrosoftChapVersionTwo,
		"unencryptedpassword":     MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls_UnencryptedPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationNonEapAuthenticationMethodForEapTtls(input)
	return &out, nil
}
