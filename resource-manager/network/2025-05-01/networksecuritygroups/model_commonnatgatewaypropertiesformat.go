package networksecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonNatGatewayPropertiesFormat struct {
	IdleTimeoutInMinutes *int64               `json:"idleTimeoutInMinutes,omitempty"`
	ProvisioningState    *ProvisioningState   `json:"provisioningState,omitempty"`
	PublicIPAddresses    *[]CommonSubResource `json:"publicIpAddresses,omitempty"`
	PublicIPAddressesV6  *[]CommonSubResource `json:"publicIpAddressesV6,omitempty"`
	PublicIPPrefixes     *[]CommonSubResource `json:"publicIpPrefixes,omitempty"`
	PublicIPPrefixesV6   *[]CommonSubResource `json:"publicIpPrefixesV6,omitempty"`
	ResourceGuid         *string              `json:"resourceGuid,omitempty"`
	ServiceGateway       *CommonSubResource   `json:"serviceGateway,omitempty"`
	SourceVirtualNetwork *CommonSubResource   `json:"sourceVirtualNetwork,omitempty"`
	Subnets              *[]CommonSubResource `json:"subnets,omitempty"`
}
