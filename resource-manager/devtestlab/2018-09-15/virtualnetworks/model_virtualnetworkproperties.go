package virtualnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworkProperties struct {
	AllowedSubnets             *[]Subnet         `json:"allowedSubnets,omitempty"`
	CreatedDate                *string           `json:"createdDate,omitempty"`
	Description                *string           `json:"description,omitempty"`
	ExternalProviderResourceId *string           `json:"externalProviderResourceId,omitempty"`
	ExternalSubnets            *[]ExternalSubnet `json:"externalSubnets,omitempty"`
	ProvisioningState          *string           `json:"provisioningState,omitempty"`
	SubnetOverrides            *[]SubnetOverride `json:"subnetOverrides,omitempty"`
	UniqueIdentifier           *string           `json:"uniqueIdentifier,omitempty"`
}
