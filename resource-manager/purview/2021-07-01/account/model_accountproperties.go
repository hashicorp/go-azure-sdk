package account

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountProperties struct {
	CloudConnectors            *CloudConnectors             `json:"cloudConnectors,omitempty"`
	CreatedAt                  *string                      `json:"createdAt,omitempty"`
	CreatedBy                  *string                      `json:"createdBy,omitempty"`
	CreatedByObjectId          *string                      `json:"createdByObjectId,omitempty"`
	Endpoints                  *AccountEndpoints            `json:"endpoints,omitempty"`
	FriendlyName               *string                      `json:"friendlyName,omitempty"`
	ManagedResourceGroupName   *string                      `json:"managedResourceGroupName,omitempty"`
	ManagedResources           *ManagedResources            `json:"managedResources,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState           `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess         `json:"publicNetworkAccess,omitempty"`
}
