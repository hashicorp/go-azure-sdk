package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel string

const (
	WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel_AnyDestinationAnySource WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel = "anyDestinationAnySource"
	WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel_None                    WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel = "none"
)

func PossibleValuesForWindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel() []string {
	return []string{
		string(WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel_AnyDestinationAnySource),
		string(WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel_None),
	}
}

func (s *WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel(input string) (*WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel, error) {
	vals := map[string]WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel{
		"anydestinationanysource": WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel_AnyDestinationAnySource,
		"none":                    WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel(input)
	return &out, nil
}
