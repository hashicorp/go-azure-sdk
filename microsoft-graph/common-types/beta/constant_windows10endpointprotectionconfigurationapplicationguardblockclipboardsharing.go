package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing string

const (
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockBoth            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "blockBoth"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockContainerToHost Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "blockContainerToHost"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockHostToContainer Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "blockHostToContainer"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockNone            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "blockNone"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_NotConfigured        Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockBoth),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockContainerToHost),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockHostToContainer),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockNone),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing(input string) (*Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing{
		"blockboth":            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockBoth,
		"blockcontainertohost": Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockContainerToHost,
		"blockhosttocontainer": Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockHostToContainer,
		"blocknone":            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_BlockNone,
		"notconfigured":        Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing(input)
	return &out, nil
}
