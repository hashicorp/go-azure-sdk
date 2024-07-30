package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleOperator string

const (
	Win32LobAppPowerShellScriptRuleOperator_Equal              Win32LobAppPowerShellScriptRuleOperator = "equal"
	Win32LobAppPowerShellScriptRuleOperator_GreaterThan        Win32LobAppPowerShellScriptRuleOperator = "greaterThan"
	Win32LobAppPowerShellScriptRuleOperator_GreaterThanOrEqual Win32LobAppPowerShellScriptRuleOperator = "greaterThanOrEqual"
	Win32LobAppPowerShellScriptRuleOperator_LessThan           Win32LobAppPowerShellScriptRuleOperator = "lessThan"
	Win32LobAppPowerShellScriptRuleOperator_LessThanOrEqual    Win32LobAppPowerShellScriptRuleOperator = "lessThanOrEqual"
	Win32LobAppPowerShellScriptRuleOperator_NotConfigured      Win32LobAppPowerShellScriptRuleOperator = "notConfigured"
	Win32LobAppPowerShellScriptRuleOperator_NotEqual           Win32LobAppPowerShellScriptRuleOperator = "notEqual"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleOperator() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleOperator_Equal),
		string(Win32LobAppPowerShellScriptRuleOperator_GreaterThan),
		string(Win32LobAppPowerShellScriptRuleOperator_GreaterThanOrEqual),
		string(Win32LobAppPowerShellScriptRuleOperator_LessThan),
		string(Win32LobAppPowerShellScriptRuleOperator_LessThanOrEqual),
		string(Win32LobAppPowerShellScriptRuleOperator_NotConfigured),
		string(Win32LobAppPowerShellScriptRuleOperator_NotEqual),
	}
}

func (s *Win32LobAppPowerShellScriptRuleOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRuleOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRuleOperator(input string) (*Win32LobAppPowerShellScriptRuleOperator, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleOperator{
		"equal":              Win32LobAppPowerShellScriptRuleOperator_Equal,
		"greaterthan":        Win32LobAppPowerShellScriptRuleOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppPowerShellScriptRuleOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppPowerShellScriptRuleOperator_LessThan,
		"lessthanorequal":    Win32LobAppPowerShellScriptRuleOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppPowerShellScriptRuleOperator_NotConfigured,
		"notequal":           Win32LobAppPowerShellScriptRuleOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleOperator(input)
	return &out, nil
}
