package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps                IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "AllApps"
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked                IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "Blocked"
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps            IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedApps"
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "ManagedAppsWithPasteIn"
)

func PossibleValuesForIosManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps),
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked),
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps),
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn),
	}
}

func parseIosManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*IosManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]IosManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                IosManagedAppProtectionAllowedOutboundClipboardSharingLevelallApps,
		"blocked":                IosManagedAppProtectionAllowedOutboundClipboardSharingLevelblocked,
		"managedapps":            IosManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedApps,
		"managedappswithpastein": IosManagedAppProtectionAllowedOutboundClipboardSharingLevelmanagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
