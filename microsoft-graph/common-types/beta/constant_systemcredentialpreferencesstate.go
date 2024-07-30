package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemCredentialPreferencesState string

const (
	SystemCredentialPreferencesState_Default  SystemCredentialPreferencesState = "default"
	SystemCredentialPreferencesState_Disabled SystemCredentialPreferencesState = "disabled"
	SystemCredentialPreferencesState_Enabled  SystemCredentialPreferencesState = "enabled"
)

func PossibleValuesForSystemCredentialPreferencesState() []string {
	return []string{
		string(SystemCredentialPreferencesState_Default),
		string(SystemCredentialPreferencesState_Disabled),
		string(SystemCredentialPreferencesState_Enabled),
	}
}

func (s *SystemCredentialPreferencesState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSystemCredentialPreferencesState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSystemCredentialPreferencesState(input string) (*SystemCredentialPreferencesState, error) {
	vals := map[string]SystemCredentialPreferencesState{
		"default":  SystemCredentialPreferencesState_Default,
		"disabled": SystemCredentialPreferencesState_Disabled,
		"enabled":  SystemCredentialPreferencesState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SystemCredentialPreferencesState(input)
	return &out, nil
}
