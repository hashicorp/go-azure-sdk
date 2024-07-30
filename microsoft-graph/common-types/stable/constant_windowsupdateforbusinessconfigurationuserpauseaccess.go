package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUserPauseAccess string

const (
	WindowsUpdateForBusinessConfigurationUserPauseAccess_Disabled      WindowsUpdateForBusinessConfigurationUserPauseAccess = "disabled"
	WindowsUpdateForBusinessConfigurationUserPauseAccess_Enabled       WindowsUpdateForBusinessConfigurationUserPauseAccess = "enabled"
	WindowsUpdateForBusinessConfigurationUserPauseAccess_NotConfigured WindowsUpdateForBusinessConfigurationUserPauseAccess = "notConfigured"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUserPauseAccess() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUserPauseAccess_Disabled),
		string(WindowsUpdateForBusinessConfigurationUserPauseAccess_Enabled),
		string(WindowsUpdateForBusinessConfigurationUserPauseAccess_NotConfigured),
	}
}

func (s *WindowsUpdateForBusinessConfigurationUserPauseAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationUserPauseAccess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationUserPauseAccess(input string) (*WindowsUpdateForBusinessConfigurationUserPauseAccess, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUserPauseAccess{
		"disabled":      WindowsUpdateForBusinessConfigurationUserPauseAccess_Disabled,
		"enabled":       WindowsUpdateForBusinessConfigurationUserPauseAccess_Enabled,
		"notconfigured": WindowsUpdateForBusinessConfigurationUserPauseAccess_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUserPauseAccess(input)
	return &out, nil
}
