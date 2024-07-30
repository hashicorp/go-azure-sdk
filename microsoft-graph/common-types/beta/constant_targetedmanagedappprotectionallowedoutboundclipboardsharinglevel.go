package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "allApps"
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "blocked"
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps            TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedApps"
	TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedAppsWithPasteIn"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps),
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked),
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps),
		string(TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn),
	}
}

func (s *TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps,
		"blocked":                TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked,
		"managedapps":            TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps,
		"managedappswithpastein": TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
