package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementOperator string

const (
	Win32LobAppPowerShellScriptRequirementOperator_Equal              Win32LobAppPowerShellScriptRequirementOperator = "equal"
	Win32LobAppPowerShellScriptRequirementOperator_GreaterThan        Win32LobAppPowerShellScriptRequirementOperator = "greaterThan"
	Win32LobAppPowerShellScriptRequirementOperator_GreaterThanOrEqual Win32LobAppPowerShellScriptRequirementOperator = "greaterThanOrEqual"
	Win32LobAppPowerShellScriptRequirementOperator_LessThan           Win32LobAppPowerShellScriptRequirementOperator = "lessThan"
	Win32LobAppPowerShellScriptRequirementOperator_LessThanOrEqual    Win32LobAppPowerShellScriptRequirementOperator = "lessThanOrEqual"
	Win32LobAppPowerShellScriptRequirementOperator_NotConfigured      Win32LobAppPowerShellScriptRequirementOperator = "notConfigured"
	Win32LobAppPowerShellScriptRequirementOperator_NotEqual           Win32LobAppPowerShellScriptRequirementOperator = "notEqual"
)

func PossibleValuesForWin32LobAppPowerShellScriptRequirementOperator() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRequirementOperator_Equal),
		string(Win32LobAppPowerShellScriptRequirementOperator_GreaterThan),
		string(Win32LobAppPowerShellScriptRequirementOperator_GreaterThanOrEqual),
		string(Win32LobAppPowerShellScriptRequirementOperator_LessThan),
		string(Win32LobAppPowerShellScriptRequirementOperator_LessThanOrEqual),
		string(Win32LobAppPowerShellScriptRequirementOperator_NotConfigured),
		string(Win32LobAppPowerShellScriptRequirementOperator_NotEqual),
	}
}

func (s *Win32LobAppPowerShellScriptRequirementOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRequirementOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRequirementOperator(input string) (*Win32LobAppPowerShellScriptRequirementOperator, error) {
	vals := map[string]Win32LobAppPowerShellScriptRequirementOperator{
		"equal":              Win32LobAppPowerShellScriptRequirementOperator_Equal,
		"greaterthan":        Win32LobAppPowerShellScriptRequirementOperator_GreaterThan,
		"greaterthanorequal": Win32LobAppPowerShellScriptRequirementOperator_GreaterThanOrEqual,
		"lessthan":           Win32LobAppPowerShellScriptRequirementOperator_LessThan,
		"lessthanorequal":    Win32LobAppPowerShellScriptRequirementOperator_LessThanOrEqual,
		"notconfigured":      Win32LobAppPowerShellScriptRequirementOperator_NotConfigured,
		"notequal":           Win32LobAppPowerShellScriptRequirementOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRequirementOperator(input)
	return &out, nil
}
