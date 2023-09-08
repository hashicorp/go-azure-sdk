package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "AllApps"
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "Blocked"
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps            TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedApps"
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedAppsWithPasteIn"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps),
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked),
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps),
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn),
	}
}

func parseTargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps,
		"blocked":                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked,
		"managedapps":            TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps,
		"managedappswithpastein": TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
