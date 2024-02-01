package workspaceazureadonlyauthentications

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StateValue string

const (
	StateValueConsistent   StateValue = "Consistent"
	StateValueInConsistent StateValue = "InConsistent"
	StateValueUpdating     StateValue = "Updating"
)

func PossibleValuesForStateValue() []string {
	return []string{
		string(StateValueConsistent),
		string(StateValueInConsistent),
		string(StateValueUpdating),
	}
}

func (s *StateValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStateValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStateValue(input string) (*StateValue, error) {
	vals := map[string]StateValue{
		"consistent":   StateValueConsistent,
		"inconsistent": StateValueInConsistent,
		"updating":     StateValueUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StateValue(input)
	return &out, nil
}
