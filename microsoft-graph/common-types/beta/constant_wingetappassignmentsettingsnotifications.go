package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppAssignmentSettingsNotifications string

const (
	WinGetAppAssignmentSettingsNotifications_HideAll    WinGetAppAssignmentSettingsNotifications = "hideAll"
	WinGetAppAssignmentSettingsNotifications_ShowAll    WinGetAppAssignmentSettingsNotifications = "showAll"
	WinGetAppAssignmentSettingsNotifications_ShowReboot WinGetAppAssignmentSettingsNotifications = "showReboot"
)

func PossibleValuesForWinGetAppAssignmentSettingsNotifications() []string {
	return []string{
		string(WinGetAppAssignmentSettingsNotifications_HideAll),
		string(WinGetAppAssignmentSettingsNotifications_ShowAll),
		string(WinGetAppAssignmentSettingsNotifications_ShowReboot),
	}
}

func (s *WinGetAppAssignmentSettingsNotifications) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWinGetAppAssignmentSettingsNotifications(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWinGetAppAssignmentSettingsNotifications(input string) (*WinGetAppAssignmentSettingsNotifications, error) {
	vals := map[string]WinGetAppAssignmentSettingsNotifications{
		"hideall":    WinGetAppAssignmentSettingsNotifications_HideAll,
		"showall":    WinGetAppAssignmentSettingsNotifications_ShowAll,
		"showreboot": WinGetAppAssignmentSettingsNotifications_ShowReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WinGetAppAssignmentSettingsNotifications(input)
	return &out, nil
}
