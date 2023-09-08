package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing string

const (
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockBoth            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "BlockBoth"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockContainerToHost Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "BlockContainerToHost"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockHostToContainer Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "BlockHostToContainer"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockNone            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "BlockNone"
	Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingnotConfigured        Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockBoth),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockContainerToHost),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockHostToContainer),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockNone),
		string(Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingnotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing(input string) (*Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing{
		"blockboth":            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockBoth,
		"blockcontainertohost": Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockContainerToHost,
		"blockhosttocontainer": Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockHostToContainer,
		"blocknone":            Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingblockNone,
		"notconfigured":        Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharingnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing(input)
	return &out, nil
}
