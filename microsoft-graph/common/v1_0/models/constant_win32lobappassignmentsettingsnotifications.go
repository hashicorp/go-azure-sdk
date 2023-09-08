package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAssignmentSettingsNotifications string

const (
	Win32LobAppAssignmentSettingsNotificationshideAll    Win32LobAppAssignmentSettingsNotifications = "HideAll"
	Win32LobAppAssignmentSettingsNotificationsshowAll    Win32LobAppAssignmentSettingsNotifications = "ShowAll"
	Win32LobAppAssignmentSettingsNotificationsshowReboot Win32LobAppAssignmentSettingsNotifications = "ShowReboot"
)

func PossibleValuesForWin32LobAppAssignmentSettingsNotifications() []string {
	return []string{
		string(Win32LobAppAssignmentSettingsNotificationshideAll),
		string(Win32LobAppAssignmentSettingsNotificationsshowAll),
		string(Win32LobAppAssignmentSettingsNotificationsshowReboot),
	}
}

func parseWin32LobAppAssignmentSettingsNotifications(input string) (*Win32LobAppAssignmentSettingsNotifications, error) {
	vals := map[string]Win32LobAppAssignmentSettingsNotifications{
		"hideall":    Win32LobAppAssignmentSettingsNotificationshideAll,
		"showall":    Win32LobAppAssignmentSettingsNotificationsshowAll,
		"showreboot": Win32LobAppAssignmentSettingsNotificationsshowReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppAssignmentSettingsNotifications(input)
	return &out, nil
}
