package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppInstallExperienceRunAsAccount string

const (
	WinGetAppInstallExperienceRunAsAccount_System WinGetAppInstallExperienceRunAsAccount = "system"
	WinGetAppInstallExperienceRunAsAccount_User   WinGetAppInstallExperienceRunAsAccount = "user"
)

func PossibleValuesForWinGetAppInstallExperienceRunAsAccount() []string {
	return []string{
		string(WinGetAppInstallExperienceRunAsAccount_System),
		string(WinGetAppInstallExperienceRunAsAccount_User),
	}
}

func (s *WinGetAppInstallExperienceRunAsAccount) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWinGetAppInstallExperienceRunAsAccount(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWinGetAppInstallExperienceRunAsAccount(input string) (*WinGetAppInstallExperienceRunAsAccount, error) {
	vals := map[string]WinGetAppInstallExperienceRunAsAccount{
		"system": WinGetAppInstallExperienceRunAsAccount_System,
		"user":   WinGetAppInstallExperienceRunAsAccount_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WinGetAppInstallExperienceRunAsAccount(input)
	return &out, nil
}
