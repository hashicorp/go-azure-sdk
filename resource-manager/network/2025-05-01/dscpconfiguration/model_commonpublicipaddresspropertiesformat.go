package dscpconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonPublicIPAddressPropertiesFormat struct {
	DdosSettings             *CommonDdosSettings               `json:"ddosSettings,omitempty"`
	DeleteOption             *DeleteOptions                    `json:"deleteOption,omitempty"`
	DnsSettings              *CommonPublicIPAddressDnsSettings `json:"dnsSettings,omitempty"`
	IPAddress                *string                           `json:"ipAddress,omitempty"`
	IPConfiguration          *CommonIPConfiguration            `json:"ipConfiguration,omitempty"`
	IPTags                   *[]CommonIPTag                    `json:"ipTags,omitempty"`
	IdleTimeoutInMinutes     *int64                            `json:"idleTimeoutInMinutes,omitempty"`
	LinkedPublicIPAddress    *CommonPublicIPAddress            `json:"linkedPublicIPAddress,omitempty"`
	MigrationPhase           *PublicIPAddressMigrationPhase    `json:"migrationPhase,omitempty"`
	NatGateway               *CommonNatGateway                 `json:"natGateway,omitempty"`
	ProvisioningState        *ProvisioningState                `json:"provisioningState,omitempty"`
	PublicIPAddressVersion   *IPVersion                        `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *IPAllocationMethod               `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *CommonSubResource                `json:"publicIPPrefix,omitempty"`
	ResourceGuid             *string                           `json:"resourceGuid,omitempty"`
	ServicePublicIPAddress   *CommonPublicIPAddress            `json:"servicePublicIPAddress,omitempty"`
}
