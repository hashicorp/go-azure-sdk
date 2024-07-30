package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodConfigurationState string

const (
	MicrosoftAuthenticatorAuthenticationMethodConfigurationState_Disabled MicrosoftAuthenticatorAuthenticationMethodConfigurationState = "disabled"
	MicrosoftAuthenticatorAuthenticationMethodConfigurationState_Enabled  MicrosoftAuthenticatorAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodConfigurationState() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodConfigurationState_Disabled),
		string(MicrosoftAuthenticatorAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *MicrosoftAuthenticatorAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftAuthenticatorAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftAuthenticatorAuthenticationMethodConfigurationState(input string) (*MicrosoftAuthenticatorAuthenticationMethodConfigurationState, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodConfigurationState{
		"disabled": MicrosoftAuthenticatorAuthenticationMethodConfigurationState_Disabled,
		"enabled":  MicrosoftAuthenticatorAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodConfigurationState(input)
	return &out, nil
}
