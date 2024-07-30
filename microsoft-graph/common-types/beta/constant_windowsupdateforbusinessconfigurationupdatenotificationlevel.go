package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUpdateNotificationLevel string

const (
	WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_DefaultNotifications    WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "defaultNotifications"
	WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_DisableAllNotifications WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "disableAllNotifications"
	WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_NotConfigured           WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "notConfigured"
	WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_RestartWarningsOnly     WindowsUpdateForBusinessConfigurationUpdateNotificationLevel = "restartWarningsOnly"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUpdateNotificationLevel() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_DefaultNotifications),
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_DisableAllNotifications),
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_NotConfigured),
		string(WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_RestartWarningsOnly),
	}
}

func (s *WindowsUpdateForBusinessConfigurationUpdateNotificationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationUpdateNotificationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationUpdateNotificationLevel(input string) (*WindowsUpdateForBusinessConfigurationUpdateNotificationLevel, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUpdateNotificationLevel{
		"defaultnotifications":    WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_DefaultNotifications,
		"disableallnotifications": WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_DisableAllNotifications,
		"notconfigured":           WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_NotConfigured,
		"restartwarningsonly":     WindowsUpdateForBusinessConfigurationUpdateNotificationLevel_RestartWarningsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUpdateNotificationLevel(input)
	return &out, nil
}
