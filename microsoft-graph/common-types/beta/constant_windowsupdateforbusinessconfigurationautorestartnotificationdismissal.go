package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal string

const (
	WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_Automatic     WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal = "automatic"
	WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_NotConfigured WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal = "notConfigured"
	WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_User          WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal = "user"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_Automatic),
		string(WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_NotConfigured),
		string(WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_User),
	}
}

func (s *WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal(input string) (*WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal{
		"automatic":     WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_Automatic,
		"notconfigured": WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_NotConfigured,
		"user":          WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal(input)
	return &out, nil
}
