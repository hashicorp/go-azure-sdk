package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	ManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps                ManagedAppProtectionAllowedOutboundClipboardSharingLevel = "allApps"
	ManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked                ManagedAppProtectionAllowedOutboundClipboardSharingLevel = "blocked"
	ManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps            ManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedApps"
	ManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn ManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedAppsWithPasteIn"
)

func PossibleValuesForManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(ManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps),
		string(ManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked),
		string(ManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps),
		string(ManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn),
	}
}

func (s *ManagedAppProtectionAllowedOutboundClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAllowedOutboundClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*ManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]ManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                ManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps,
		"blocked":                ManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked,
		"managedapps":            ManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps,
		"managedappswithpastein": ManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
