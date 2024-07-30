package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemporaryAccessPassAuthenticationMethodConfigurationState string

const (
	TemporaryAccessPassAuthenticationMethodConfigurationState_Disabled TemporaryAccessPassAuthenticationMethodConfigurationState = "disabled"
	TemporaryAccessPassAuthenticationMethodConfigurationState_Enabled  TemporaryAccessPassAuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForTemporaryAccessPassAuthenticationMethodConfigurationState() []string {
	return []string{
		string(TemporaryAccessPassAuthenticationMethodConfigurationState_Disabled),
		string(TemporaryAccessPassAuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *TemporaryAccessPassAuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTemporaryAccessPassAuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTemporaryAccessPassAuthenticationMethodConfigurationState(input string) (*TemporaryAccessPassAuthenticationMethodConfigurationState, error) {
	vals := map[string]TemporaryAccessPassAuthenticationMethodConfigurationState{
		"disabled": TemporaryAccessPassAuthenticationMethodConfigurationState_Disabled,
		"enabled":  TemporaryAccessPassAuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TemporaryAccessPassAuthenticationMethodConfigurationState(input)
	return &out, nil
}
