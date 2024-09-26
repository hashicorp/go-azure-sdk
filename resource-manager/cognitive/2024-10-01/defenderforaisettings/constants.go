package defenderforaisettings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderForAISettingState string

const (
	DefenderForAISettingStateDisabled DefenderForAISettingState = "Disabled"
	DefenderForAISettingStateEnabled  DefenderForAISettingState = "Enabled"
)

func PossibleValuesForDefenderForAISettingState() []string {
	return []string{
		string(DefenderForAISettingStateDisabled),
		string(DefenderForAISettingStateEnabled),
	}
}

func (s *DefenderForAISettingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderForAISettingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderForAISettingState(input string) (*DefenderForAISettingState, error) {
	vals := map[string]DefenderForAISettingState{
		"disabled": DefenderForAISettingStateDisabled,
		"enabled":  DefenderForAISettingStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderForAISettingState(input)
	return &out, nil
}
