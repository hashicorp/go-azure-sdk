package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAssignmentSettingsDeliveryOptimizationPriority string

const (
	Win32LobAppAssignmentSettingsDeliveryOptimizationPriority_Foreground    Win32LobAppAssignmentSettingsDeliveryOptimizationPriority = "foreground"
	Win32LobAppAssignmentSettingsDeliveryOptimizationPriority_NotConfigured Win32LobAppAssignmentSettingsDeliveryOptimizationPriority = "notConfigured"
)

func PossibleValuesForWin32LobAppAssignmentSettingsDeliveryOptimizationPriority() []string {
	return []string{
		string(Win32LobAppAssignmentSettingsDeliveryOptimizationPriority_Foreground),
		string(Win32LobAppAssignmentSettingsDeliveryOptimizationPriority_NotConfigured),
	}
}

func (s *Win32LobAppAssignmentSettingsDeliveryOptimizationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppAssignmentSettingsDeliveryOptimizationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppAssignmentSettingsDeliveryOptimizationPriority(input string) (*Win32LobAppAssignmentSettingsDeliveryOptimizationPriority, error) {
	vals := map[string]Win32LobAppAssignmentSettingsDeliveryOptimizationPriority{
		"foreground":    Win32LobAppAssignmentSettingsDeliveryOptimizationPriority_Foreground,
		"notconfigured": Win32LobAppAssignmentSettingsDeliveryOptimizationPriority_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppAssignmentSettingsDeliveryOptimizationPriority(input)
	return &out, nil
}
