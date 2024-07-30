package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppFileSystemRuleOperator string

const (
	Win32LobAppFileSystemRuleOperator_Equal              Win32LobAppFileSystemRuleOperator = "equal"
	Win32LobAppFileSystemRuleOperator_GreaterThan        Win32LobAppFileSystemRuleOperator = "greaterThan"
	Win32LobAppFileSystemRuleOperator_GreaterThanOrEqual Win32LobAppFileSystemRuleOperator = "greaterThanOrEqual"
	Win32LobAppFileSystemRuleOperator_LessThan           Win32LobAppFileSystemRuleOperator = "lessThan"
	Win32LobAppFileSystemRuleOperator_LessThanOrEqual    Win32LobAppFileSystemRuleOperator = "lessThanOrEqual"
	Win32LobAppFileSystemRuleOperator_NotConfigured      Win32LobAppFileSystemRuleOperator = "notConfigured"
	Win32LobAppFileSystemRuleOperator_NotEqual           Win32LobAppFileSystemRuleOperator = "notEqual"
)

func PossibleValuesForWin32LobAppFileSystemRuleOperator() []string {
	return []string{
		string(Win32LobAppFileSystemRuleOperator_Equal),
		string(Win32LobAppFileSystemRuleOperator_GreaterThan),
		string(Win32LobAppFileSystemRuleOperator_GreaterThanOrEqual),
		string(Win32LobAppFileSystemRuleOperator_LessThan),
		string(Win32LobAppFileSystemRuleOperator_LessThanOrEqual),
		string(Win32LobAppFileSystemRuleOperator_NotConfigured),
		string(Win32LobAppFileSystemRuleOperator_NotEqual),
	}
}

func (s *Win32LobAppFileSystemRuleOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppFileSystemRuleOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppFileSystemRuleOperator(input string) (*Win32LobAppFileSystemRuleOperator, error) {
	vals := map[string]Win32LobAppFileSystemRuleOperator{
		"equal":              Win32LobAppFileSystemRuleOperator_Equal,
		"greaterthan":        Win32LobAppFileSystemRuleOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppFileSystemRuleOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppFileSystemRuleOperator_LessThan,
		"lessthanorequal":    Win32LobAppFileSystemRuleOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppFileSystemRuleOperator_NotConfigured,
		"notequal":           Win32LobAppFileSystemRuleOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppFileSystemRuleOperator(input)
	return &out, nil
}
