package digitaltwinsinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DigitalTwinsProperties struct {
	CreatedTime                *string                      `json:"createdTime,omitempty"`
	HostName                   *string                      `json:"hostName,omitempty"`
	LastUpdatedTime            *string                      `json:"lastUpdatedTime,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState           `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess         `json:"publicNetworkAccess,omitempty"`
}
