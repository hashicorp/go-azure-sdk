package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus string

const (
	NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatusdisabled NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus = "Disabled"
	NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatusenabled  NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus = "Enabled"
)

func PossibleValuesForNetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus() []string {
	return []string{
		string(NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatusdisabled),
		string(NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatusenabled),
	}
}

func parseNetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus(input string) (*NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus, error) {
	vals := map[string]NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus{
		"disabled": NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatusdisabled,
		"enabled":  NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatusenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessCrossTenantAccessSettingsNetworkPacketTaggingStatus(input)
	return &out, nil
}
