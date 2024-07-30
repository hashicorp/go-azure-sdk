package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAutoUpdateSupersededApps string

const (
	Win32LobAppAutoUpdateSupersededApps_Enabled       Win32LobAppAutoUpdateSupersededApps = "enabled"
	Win32LobAppAutoUpdateSupersededApps_NotConfigured Win32LobAppAutoUpdateSupersededApps = "notConfigured"
)

func PossibleValuesForWin32LobAppAutoUpdateSupersededApps() []string {
	return []string{
		string(Win32LobAppAutoUpdateSupersededApps_Enabled),
		string(Win32LobAppAutoUpdateSupersededApps_NotConfigured),
	}
}

func (s *Win32LobAppAutoUpdateSupersededApps) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppAutoUpdateSupersededApps(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppAutoUpdateSupersededApps(input string) (*Win32LobAppAutoUpdateSupersededApps, error) {
	vals := map[string]Win32LobAppAutoUpdateSupersededApps{
		"enabled":       Win32LobAppAutoUpdateSupersededApps_Enabled,
		"notconfigured": Win32LobAppAutoUpdateSupersededApps_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppAutoUpdateSupersededApps(input)
	return &out, nil
}
