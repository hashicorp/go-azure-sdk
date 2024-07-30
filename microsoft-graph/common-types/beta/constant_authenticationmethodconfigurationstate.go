package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodConfigurationState string

const (
	AuthenticationMethodConfigurationState_Disabled AuthenticationMethodConfigurationState = "disabled"
	AuthenticationMethodConfigurationState_Enabled  AuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForAuthenticationMethodConfigurationState() []string {
	return []string{
		string(AuthenticationMethodConfigurationState_Disabled),
		string(AuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *AuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodConfigurationState(input string) (*AuthenticationMethodConfigurationState, error) {
	vals := map[string]AuthenticationMethodConfigurationState{
		"disabled": AuthenticationMethodConfigurationState_Disabled,
		"enabled":  AuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodConfigurationState(input)
	return &out, nil
}
