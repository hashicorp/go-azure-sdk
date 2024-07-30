package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuiltInIdentityProviderState string

const (
	BuiltInIdentityProviderState_Disabled BuiltInIdentityProviderState = "disabled"
	BuiltInIdentityProviderState_Enabled  BuiltInIdentityProviderState = "enabled"
)

func PossibleValuesForBuiltInIdentityProviderState() []string {
	return []string{
		string(BuiltInIdentityProviderState_Disabled),
		string(BuiltInIdentityProviderState_Enabled),
	}
}

func (s *BuiltInIdentityProviderState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBuiltInIdentityProviderState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBuiltInIdentityProviderState(input string) (*BuiltInIdentityProviderState, error) {
	vals := map[string]BuiltInIdentityProviderState{
		"disabled": BuiltInIdentityProviderState_Disabled,
		"enabled":  BuiltInIdentityProviderState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuiltInIdentityProviderState(input)
	return &out, nil
}
