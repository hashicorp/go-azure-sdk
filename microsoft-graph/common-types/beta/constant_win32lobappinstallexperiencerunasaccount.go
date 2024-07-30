package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppInstallExperienceRunAsAccount string

const (
	Win32LobAppInstallExperienceRunAsAccount_System Win32LobAppInstallExperienceRunAsAccount = "system"
	Win32LobAppInstallExperienceRunAsAccount_User   Win32LobAppInstallExperienceRunAsAccount = "user"
)

func PossibleValuesForWin32LobAppInstallExperienceRunAsAccount() []string {
	return []string{
		string(Win32LobAppInstallExperienceRunAsAccount_System),
		string(Win32LobAppInstallExperienceRunAsAccount_User),
	}
}

func (s *Win32LobAppInstallExperienceRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppInstallExperienceRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppInstallExperienceRunAsAccount(input string) (*Win32LobAppInstallExperienceRunAsAccount, error) {
	vals := map[string]Win32LobAppInstallExperienceRunAsAccount{
		"system": Win32LobAppInstallExperienceRunAsAccount_System,
		"user":   Win32LobAppInstallExperienceRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppInstallExperienceRunAsAccount(input)
	return &out, nil
}
