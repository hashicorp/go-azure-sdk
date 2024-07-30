package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppPowerShellScriptRequirementRunAsAccount string

const (
	Win32LobAppPowerShellScriptRequirementRunAsAccount_System Win32LobAppPowerShellScriptRequirementRunAsAccount = "system"
	Win32LobAppPowerShellScriptRequirementRunAsAccount_User   Win32LobAppPowerShellScriptRequirementRunAsAccount = "user"
)

func PossibleValuesForWin32LobAppPowerShellScriptRequirementRunAsAccount() []string {
	return []string{
		string(Win32LobAppPowerShellScriptRequirementRunAsAccount_System),
		string(Win32LobAppPowerShellScriptRequirementRunAsAccount_User),
	}
}

func (s *Win32LobAppPowerShellScriptRequirementRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppPowerShellScriptRequirementRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppPowerShellScriptRequirementRunAsAccount(input string) (*Win32LobAppPowerShellScriptRequirementRunAsAccount, error) {
	vals := map[string]Win32LobAppPowerShellScriptRequirementRunAsAccount{
		"system": Win32LobAppPowerShellScriptRequirementRunAsAccount_System,
		"user":   Win32LobAppPowerShellScriptRequirementRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppPowerShellScriptRequirementRunAsAccount(input)
	return &out, nil
}
