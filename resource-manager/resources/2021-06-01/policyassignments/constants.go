package policyassignments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnforcementMode string

const (
	EnforcementModeDefault      EnforcementMode = "Default"
	EnforcementModeDoNotEnforce EnforcementMode = "DoNotEnforce"
)

func PossibleValuesForEnforcementMode() []string {
	return []string{
		string(EnforcementModeDefault),
		string(EnforcementModeDoNotEnforce),
	}
}

func (s *EnforcementMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnforcementMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnforcementMode(input string) (*EnforcementMode, error) {
	vals := map[string]EnforcementMode{
		"default":      EnforcementModeDefault,
		"donotenforce": EnforcementModeDoNotEnforce,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnforcementMode(input)
	return &out, nil
}
