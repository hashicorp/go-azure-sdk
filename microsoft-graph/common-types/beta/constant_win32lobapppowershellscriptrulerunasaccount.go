package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRuleRunAsAccount string

const (
	Win32LobAppPowerShellScriptRuleRunAsAccount_System Win32LobAppPowerShellScriptRuleRunAsAccount = "system"
	Win32LobAppPowerShellScriptRuleRunAsAccount_User   Win32LobAppPowerShellScriptRuleRunAsAccount = "user"
)

func PossibleValuesForWin32LobAppPowerShellScriptRuleRunAsAccount() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRuleRunAsAccount_System),
		string(Win32LobAppPowerShellScriptRuleRunAsAccount_User),
	}
}

func (s *Win32LobAppPowerShellScriptRuleRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRuleRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRuleRunAsAccount(input string) (*Win32LobAppPowerShellScriptRuleRunAsAccount, error) {
	vals := map[string]Win32LobAppPowerShellScriptRuleRunAsAccount{
		"system": Win32LobAppPowerShellScriptRuleRunAsAccount_System,
		"user":   Win32LobAppPowerShellScriptRuleRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRuleRunAsAccount(input)
	return &out, nil
}
