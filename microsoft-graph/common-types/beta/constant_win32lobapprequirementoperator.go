package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRequirementOperator string

const (
	Win32LobAppRequirementOperator_Equal              Win32LobAppRequirementOperator = "equal"
	Win32LobAppRequirementOperator_GreaterThan        Win32LobAppRequirementOperator = "greaterThan"
	Win32LobAppRequirementOperator_GreaterThanOrEqual Win32LobAppRequirementOperator = "greaterThanOrEqual"
	Win32LobAppRequirementOperator_LessThan           Win32LobAppRequirementOperator = "lessThan"
	Win32LobAppRequirementOperator_LessThanOrEqual    Win32LobAppRequirementOperator = "lessThanOrEqual"
	Win32LobAppRequirementOperator_NotConfigured      Win32LobAppRequirementOperator = "notConfigured"
	Win32LobAppRequirementOperator_NotEqual           Win32LobAppRequirementOperator = "notEqual"
)

func PossibleValuesForWin32LobAppRequirementOperator() []string {
	return []string{
		string(Win32LobAppRequirementOperator_Equal),
		string(Win32LobAppRequirementOperator_GreaterThan),
		string(Win32LobAppRequirementOperator_GreaterThanOrEqual),
		string(Win32LobAppRequirementOperator_LessThan),
		string(Win32LobAppRequirementOperator_LessThanOrEqual),
		string(Win32LobAppRequirementOperator_NotConfigured),
		string(Win32LobAppRequirementOperator_NotEqual),
	}
}

func (s *Win32LobAppRequirementOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRequirementOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRequirementOperator(input string) (*Win32LobAppRequirementOperator, error) {
	vals := map[string]Win32LobAppRequirementOperator{
		"equal":              Win32LobAppRequirementOperator_Equal,
		"greaterthan":        Win32LobAppRequirementOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppRequirementOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppRequirementOperator_LessThan,
		"lessthanorequal":    Win32LobAppRequirementOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppRequirementOperator_NotConfigured,
		"notequal":           Win32LobAppRequirementOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRequirementOperator(input)
	return &out, nil
}
