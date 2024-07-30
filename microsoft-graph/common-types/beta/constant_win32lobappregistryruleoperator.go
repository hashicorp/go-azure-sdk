package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRuleOperator string

const (
	Win32LobAppRegistryRuleOperator_Equal              Win32LobAppRegistryRuleOperator = "equal"
	Win32LobAppRegistryRuleOperator_GreaterThan        Win32LobAppRegistryRuleOperator = "greaterThan"
	Win32LobAppRegistryRuleOperator_GreaterThanOrEqual Win32LobAppRegistryRuleOperator = "greaterThanOrEqual"
	Win32LobAppRegistryRuleOperator_LessThan           Win32LobAppRegistryRuleOperator = "lessThan"
	Win32LobAppRegistryRuleOperator_LessThanOrEqual    Win32LobAppRegistryRuleOperator = "lessThanOrEqual"
	Win32LobAppRegistryRuleOperator_NotConfigured      Win32LobAppRegistryRuleOperator = "notConfigured"
	Win32LobAppRegistryRuleOperator_NotEqual           Win32LobAppRegistryRuleOperator = "notEqual"
)

func PossibleValuesForWin32LobAppRegistryRuleOperator() []string {
	return []string{
		string(Win32LobAppRegistryRuleOperator_Equal),
		string(Win32LobAppRegistryRuleOperator_GreaterThan),
		string(Win32LobAppRegistryRuleOperator_GreaterThanOrEqual),
		string(Win32LobAppRegistryRuleOperator_LessThan),
		string(Win32LobAppRegistryRuleOperator_LessThanOrEqual),
		string(Win32LobAppRegistryRuleOperator_NotConfigured),
		string(Win32LobAppRegistryRuleOperator_NotEqual),
	}
}

func (s *Win32LobAppRegistryRuleOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppRegistryRuleOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppRegistryRuleOperator(input string) (*Win32LobAppRegistryRuleOperator, error) {
	vals := map[string]Win32LobAppRegistryRuleOperator{
		"equal":              Win32LobAppRegistryRuleOperator_Equal,
		"greaterthan":        Win32LobAppRegistryRuleOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppRegistryRuleOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppRegistryRuleOperator_LessThan,
		"lessthanorequal":    Win32LobAppRegistryRuleOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppRegistryRuleOperator_NotConfigured,
		"notequal":           Win32LobAppRegistryRuleOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppRegistryRuleOperator(input)
	return &out, nil
}
