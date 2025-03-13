package workspacemanagerconfigurations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Mode string

const (
	ModeDisabled Mode = "Disabled"
	ModeEnabled  Mode = "Enabled"
)

func PossibleValuesForMode() []string {
	return []string{
		string(ModeDisabled),
		string(ModeEnabled),
	}
}

func (s *Mode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMode(input string) (*Mode, error) {
	vals := map[string]Mode{
		"disabled": ModeDisabled,
		"enabled":  ModeEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Mode(input)
	return &out, nil
}
