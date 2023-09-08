package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "AllApps"
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "Blocked"
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps            DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedApps"
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedAppsWithPasteIn"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps),
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked),
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps),
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn),
	}
}

func parseDefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps,
		"blocked":                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked,
		"managedapps":            DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps,
		"managedappswithpastein": DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
