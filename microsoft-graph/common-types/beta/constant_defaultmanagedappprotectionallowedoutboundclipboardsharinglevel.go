package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "allApps"
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "blocked"
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps            DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedApps"
	DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedAppsWithPasteIn"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps),
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked),
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps),
		string(DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn),
	}
}

func (s *DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps,
		"blocked":                DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked,
		"managedapps":            DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps,
		"managedappswithpastein": DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
