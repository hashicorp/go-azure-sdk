package account

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountProperties struct {
	AccountStatus                       *AccountStatus                       `json:"accountStatus,omitempty"`
	CloudConnectors                     *CloudConnectors                     `json:"cloudConnectors,omitempty"`
	CreatedAt                           *string                              `json:"createdAt,omitempty"`
	CreatedBy                           *string                              `json:"createdBy,omitempty"`
	CreatedByObjectId                   *string                              `json:"createdByObjectId,omitempty"`
	Endpoints                           *AccountEndpoints                    `json:"endpoints,omitempty"`
	FriendlyName                        *string                              `json:"friendlyName,omitempty"`
	ManagedEventHubState                *ManagedEventHubState                `json:"managedEventHubState,omitempty"`
	ManagedResourceGroupName            *string                              `json:"managedResourceGroupName,omitempty"`
	ManagedResources                    *ManagedResources                    `json:"managedResources,omitempty"`
	ManagedResourcesPublicNetworkAccess *ManagedResourcesPublicNetworkAccess `json:"managedResourcesPublicNetworkAccess,omitempty"`
	PrivateEndpointConnections          *[]PrivateEndpointConnection         `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                   *ProvisioningState                   `json:"provisioningState,omitempty"`
	PublicNetworkAccess                 *PublicNetworkAccess                 `json:"publicNetworkAccess,omitempty"`
}
