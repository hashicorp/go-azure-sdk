package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "allApps"
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "blocked"
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps            AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedApps"
	AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedAppsWithPasteIn"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps),
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked),
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps),
		string(AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn),
	}
}

func (s *AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps,
		"blocked":                AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked,
		"managedapps":            AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps,
		"managedappswithpastein": AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
