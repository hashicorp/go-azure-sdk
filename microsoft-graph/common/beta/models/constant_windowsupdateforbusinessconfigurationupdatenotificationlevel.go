package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUpdateNotificationLevel string

const (
	WindowsUpdateForBusinessConfigurationUpdateNotificationLeveldefaultNotifications    WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "DefaultNotifications"
	WindowsUpdateForBusinessConfigurationUpdateNotificationLeveldisableAllNotifications WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "DisableAllNotifications"
	WindowsUpdateForBusinessConfigurationUpdateNotificationLevelnotConfigured           WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "NotConfigured"
	WindowsUpdateForBusinessConfigurationUpdateNotificationLevelrestartWarningsOnly     WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "RestartWarningsOnly"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUpdateNotificationLevel() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLeveldefaultNotifications),
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLeveldisableAllNotifications),
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLevelnotConfigured),
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLevelrestartWarningsOnly),
	}
}

func parseWindowsUpdateForBusinessConfigurationUpdateNotificationLevel(input string) (*WindowsUpdateForBusinessConfigurationUpdateNotificationLevel, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUpdateNotificationLevel{
		"defaultnotifications":    WindowsUpdateForBusinessConfigurationUpdateNotificationLeveldefaultNotifications,
		"disableallnotifications": WindowsUpdateForBusinessConfigurationUpdateNotificationLeveldisableAllNotifications,
		"notconfigured":           WindowsUpdateForBusinessConfigurationUpdateNotificationLevelnotConfigured,
		"restartwarningsonly":     WindowsUpdateForBusinessConfigurationUpdateNotificationLevelrestartWarningsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUpdateNotificationLevel(input)
	return &out, nil
}
