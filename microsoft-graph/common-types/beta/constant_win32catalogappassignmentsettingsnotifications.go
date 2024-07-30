package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32CatalogAppAssignmentSettingsNotifications string

const (
	Win32CatalogAppAssignmentSettingsNotifications_HideAll    Win32CatalogAppAssignmentSettingsNotifications = "hideAll"
	Win32CatalogAppAssignmentSettingsNotifications_ShowAll    Win32CatalogAppAssignmentSettingsNotifications = "showAll"
	Win32CatalogAppAssignmentSettingsNotifications_ShowReboot Win32CatalogAppAssignmentSettingsNotifications = "showReboot"
)

func PossibleValuesForWin32CatalogAppAssignmentSettingsNotifications() []string {
	return []string{
		string(Win32CatalogAppAssignmentSettingsNotifications_HideAll),
		string(Win32CatalogAppAssignmentSettingsNotifications_ShowAll),
		string(Win32CatalogAppAssignmentSettingsNotifications_ShowReboot),
	}
}

func (s *Win32CatalogAppAssignmentSettingsNotifications) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32CatalogAppAssignmentSettingsNotifications(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32CatalogAppAssignmentSettingsNotifications(input string) (*Win32CatalogAppAssignmentSettingsNotifications, error) {
	vals := map[string]Win32CatalogAppAssignmentSettingsNotifications{
		"hideall":    Win32CatalogAppAssignmentSettingsNotifications_HideAll,
		"showall":    Win32CatalogAppAssignmentSettingsNotifications_ShowAll,
		"showreboot": Win32CatalogAppAssignmentSettingsNotifications_ShowReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32CatalogAppAssignmentSettingsNotifications(input)
	return &out, nil
}
