package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRequirementOperator string

const (
	Win32LobAppRegistryRequirementOperator_Equal              Win32LobAppRegistryRequirementOperator = "equal"
	Win32LobAppRegistryRequirementOperator_GreaterThan        Win32LobAppRegistryRequirementOperator = "greaterThan"
	Win32LobAppRegistryRequirementOperator_GreaterThanOrEqual Win32LobAppRegistryRequirementOperator = "greaterThanOrEqual"
	Win32LobAppRegistryRequirementOperator_LessThan           Win32LobAppRegistryRequirementOperator = "lessThan"
	Win32LobAppRegistryRequirementOperator_LessThanOrEqual    Win32LobAppRegistryRequirementOperator = "lessThanOrEqual"
	Win32LobAppRegistryRequirementOperator_NotConfigured      Win32LobAppRegistryRequirementOperator = "notConfigured"
	Win32LobAppRegistryRequirementOperator_NotEqual           Win32LobAppRegistryRequirementOperator = "notEqual"
)

func PossibleValuesForWin32LobAppRegistryRequirementOperator() []string {
	return []string{
		string(Win32LobAppRegistryRequirementOperator_Equal),
		string(Win32LobAppRegistryRequirementOperator_GreaterThan),
		string(Win32LobAppRegistryRequirementOperator_GreaterThanOrEqual),
		string(Win32LobAppRegistryRequirementOperator_LessThan),
		string(Win32LobAppRegistryRequirementOperator_LessThanOrEqual),
		string(Win32LobAppRegistryRequirementOperator_NotConfigured),
		string(Win32LobAppRegistryRequirementOperator_NotEqual),
	}
}

func (s *Win32LobAppRegistryRequirementOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryRequirementOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryRequirementOperator(input string) (*Win32LobAppRegistryRequirementOperator, error) {
	vals := map[string]Win32LobAppRegistryRequirementOperator{
		"equal":              Win32LobAppRegistryRequirementOperator_Equal,
		"greaterthan":        Win32LobAppRegistryRequirementOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppRegistryRequirementOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppRegistryRequirementOperator_LessThan,
		"lessthanorequal":    Win32LobAppRegistryRequirementOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppRegistryRequirementOperator_NotConfigured,
		"notequal":           Win32LobAppRegistryRequirementOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRequirementOperator(input)
	return &out, nil
}
