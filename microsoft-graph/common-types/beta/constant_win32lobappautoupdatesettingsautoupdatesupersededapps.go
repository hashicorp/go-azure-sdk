package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps string

const (
	Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps_Enabled       Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps = "enabled"
	Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps_NotConfigured Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps = "notConfigured"
)

func PossibleValuesForWin32LobAppAutoUpdateSettingsAutoUpdateSupersededApps() []string {
	return []string{
		string(Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps_Enabled),
		string(Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps_NotConfigured),
	}
}

func (s *Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppAutoUpdateSettingsAutoUpdateSupersededApps(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppAutoUpdateSettingsAutoUpdateSupersededApps(input string) (*Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps, error) {
	vals := map[string]Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps{
		"enabled":       Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps_Enabled,
		"notconfigured": Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppAutoUpdateSettingsAutoUpdateSupersededApps(input)
	return &out, nil
}
