package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRedundancyConfiguration struct {
	ODataType          *string                                             `json:"@odata.type,omitempty"`
	RedundancyTier     *NetworkaccessRedundancyConfigurationRedundancyTier `json:"redundancyTier,omitempty"`
	ZoneLocalIpAddress *string                                             `json:"zoneLocalIpAddress,omitempty"`
}
