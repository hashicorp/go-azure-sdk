package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAssignmentSettingsNotifications string

const (
	Win32LobAppAssignmentSettingsNotifications_HideAll    Win32LobAppAssignmentSettingsNotifications = "hideAll"
	Win32LobAppAssignmentSettingsNotifications_ShowAll    Win32LobAppAssignmentSettingsNotifications = "showAll"
	Win32LobAppAssignmentSettingsNotifications_ShowReboot Win32LobAppAssignmentSettingsNotifications = "showReboot"
)

func PossibleValuesForWin32LobAppAssignmentSettingsNotifications() []string {
	return []string{
		string(Win32LobAppAssignmentSettingsNotifications_HideAll),
		string(Win32LobAppAssignmentSettingsNotifications_ShowAll),
		string(Win32LobAppAssignmentSettingsNotifications_ShowReboot),
	}
}

func (s *Win32LobAppAssignmentSettingsNotifications) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppAssignmentSettingsNotifications(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppAssignmentSettingsNotifications(input string) (*Win32LobAppAssignmentSettingsNotifications, error) {
	vals := map[string]Win32LobAppAssignmentSettingsNotifications{
		"hideall":    Win32LobAppAssignmentSettingsNotifications_HideAll,
		"showall":    Win32LobAppAssignmentSettingsNotifications_ShowAll,
		"showreboot": Win32LobAppAssignmentSettingsNotifications_ShowReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppAssignmentSettingsNotifications(input)
	return &out, nil
}
