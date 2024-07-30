package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps                IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "allApps"
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked                IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "blocked"
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps            IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedApps"
	IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn IosManagedAppProtectionAllowedOutboundClipboardSharingLevel = "managedAppsWithPasteIn"
)

func PossibleValuesForIosManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps),
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked),
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps),
		string(IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn),
	}
}

func (s *IosManagedAppProtectionAllowedOutboundClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAllowedOutboundClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*IosManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]IosManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"allapps":                IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_AllApps,
		"blocked":                IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_Blocked,
		"managedapps":            IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedApps,
		"managedappswithpastein": IosManagedAppProtectionAllowedOutboundClipboardSharingLevel_ManagedAppsWithPasteIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
