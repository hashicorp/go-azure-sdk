package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "AllApps"
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "Blocked"
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps            AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedApps"
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedAppsWithPasteIn"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps),
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked),
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps),
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn),
	}
}

func parseAndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps,
		"blocked":                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked,
		"managedapps":            AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps,
		"managedappswithpastein": AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
