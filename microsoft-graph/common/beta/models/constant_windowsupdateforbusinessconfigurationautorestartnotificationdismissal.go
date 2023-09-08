package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal string

const (
	WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissalautomatic     WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal = "Automatic"
	WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissalnotConfigured WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal = "NotConfigured"
	WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissaluser          WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal = "User"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissalautomatic),
		string(WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissalnotConfigured),
		string(WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissaluser),
	}
}

func parseWindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal(input string) (*WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal{
		"automatic":     WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissalautomatic,
		"notconfigured": WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissalnotConfigured,
		"user":          WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissaluser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal(input)
	return &out, nil
}
