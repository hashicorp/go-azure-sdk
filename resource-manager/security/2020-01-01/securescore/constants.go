package securescore

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlType string

const (
	ControlTypeBuiltIn ControlType = "BuiltIn"
	ControlTypeCustom  ControlType = "Custom"
)

func PossibleValuesForControlType() []string {
	return []string{
		string(ControlTypeBuiltIn),
		string(ControlTypeCustom),
	}
}

func (s *ControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseControlType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseControlType(input string) (*ControlType, error) {
	vals := map[string]ControlType{
		"builtin": ControlTypeBuiltIn,
		"custom":  ControlTypeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ControlType(input)
	return &out, nil
}

type ExpandControlsEnum string

const (
	ExpandControlsEnumDefinition ExpandControlsEnum = "definition"
)

func PossibleValuesForExpandControlsEnum() []string {
	return []string{
		string(ExpandControlsEnumDefinition),
	}
}

func (s *ExpandControlsEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExpandControlsEnum(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExpandControlsEnum(input string) (*ExpandControlsEnum, error) {
	vals := map[string]ExpandControlsEnum{
		"definition": ExpandControlsEnumDefinition,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpandControlsEnum(input)
	return &out, nil
}
