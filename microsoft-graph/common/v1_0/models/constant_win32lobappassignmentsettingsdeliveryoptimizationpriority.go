package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAssignmentSettingsDeliveryOptimizationPriority string

const (
	Win32LobAppAssignmentSettingsDeliveryOptimizationPriorityforeground    Win32LobAppAssignmentSettingsDeliveryOptimizationPriority = "Foreground"
	Win32LobAppAssignmentSettingsDeliveryOptimizationPrioritynotConfigured Win32LobAppAssignmentSettingsDeliveryOptimizationPriority = "NotConfigured"
)

func PossibleValuesForWin32LobAppAssignmentSettingsDeliveryOptimizationPriority() []string {
	return []string{
		string(Win32LobAppAssignmentSettingsDeliveryOptimizationPriorityforeground),
		string(Win32LobAppAssignmentSettingsDeliveryOptimizationPrioritynotConfigured),
	}
}

func parseWin32LobAppAssignmentSettingsDeliveryOptimizationPriority(input string) (*Win32LobAppAssignmentSettingsDeliveryOptimizationPriority, error) {
	vals := map[string]Win32LobAppAssignmentSettingsDeliveryOptimizationPriority{
		"foreground":    Win32LobAppAssignmentSettingsDeliveryOptimizationPriorityforeground,
		"notconfigured": Win32LobAppAssignmentSettingsDeliveryOptimizationPrioritynotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppAssignmentSettingsDeliveryOptimizationPriority(input)
	return &out, nil
}
