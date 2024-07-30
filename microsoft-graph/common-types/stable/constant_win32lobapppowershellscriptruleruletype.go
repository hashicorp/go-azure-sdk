package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleRuleType string

const (
	Win32LobAppPowerShellScriptRuleRuleType_Detection   Win32LobAppPowerShellScriptRuleRuleType = "detection"
	Win32LobAppPowerShellScriptRuleRuleType_Requirement Win32LobAppPowerShellScriptRuleRuleType = "requirement"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleRuleType() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleRuleType_Detection),
		string(Win32LobAppPowerShellScriptRuleRuleType_Requirement),
	}
}

func (s *Win32LobAppPowerShellScriptRuleRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRuleRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRuleRuleType(input string) (*Win32LobAppPowerShellScriptRuleRuleType, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleRuleType{
		"detection":   Win32LobAppPowerShellScriptRuleRuleType_Detection,
		"requirement": Win32LobAppPowerShellScriptRuleRuleType_Requirement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleRuleType(input)
	return &out, nil
}
