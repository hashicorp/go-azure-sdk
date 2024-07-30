package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodConfigurationState string

const (
	Fido2AuthenticationMethodConfigurationState_Disabled Fido2AuthenticationMethodConfigurationState = "disabled"
	Fido2AuthenticationMethodConfigurationState_Enabled  Fido2AuthenticationMethodConfigurationState = "enabled"
)

func PossibleValuesForFido2AuthenticationMethodConfigurationState() []string {
	return []string{
		string(Fido2AuthenticationMethodConfigurationState_Disabled),
		string(Fido2AuthenticationMethodConfigurationState_Enabled),
	}
}

func (s *Fido2AuthenticationMethodConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFido2AuthenticationMethodConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFido2AuthenticationMethodConfigurationState(input string) (*Fido2AuthenticationMethodConfigurationState, error) {
	vals := map[string]Fido2AuthenticationMethodConfigurationState{
		"disabled": Fido2AuthenticationMethodConfigurationState_Disabled,
		"enabled":  Fido2AuthenticationMethodConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2AuthenticationMethodConfigurationState(input)
	return &out, nil
}
