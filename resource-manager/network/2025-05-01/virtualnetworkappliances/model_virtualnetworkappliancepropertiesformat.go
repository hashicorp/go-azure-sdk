package virtualnetworkappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkAppliancePropertiesFormat struct {
	BandwidthInGbps   *string                                   `json:"bandwidthInGbps,omitempty"`
	IPConfigurations  *[]VirtualNetworkApplianceIPConfiguration `json:"ipConfigurations,omitempty"`
	ProvisioningState *ProvisioningState                        `json:"provisioningState,omitempty"`
	ResourceGuid      *string                                   `json:"resourceGuid,omitempty"`
	Subnet            *CommonSubnet                             `json:"subnet,omitempty"`
}
