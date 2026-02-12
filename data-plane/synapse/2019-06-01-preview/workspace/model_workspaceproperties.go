package workspace

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceProperties struct {
	AdlaResourceId                   *string                           `json:"adlaResourceId,omitempty"`
	ConnectivityEndpoints            *map[string]string                `json:"connectivityEndpoints,omitempty"`
	DefaultDataLakeStorage           *DataLakeStorageAccountDetails    `json:"defaultDataLakeStorage,omitempty"`
	Encryption                       *EncryptionDetails                `json:"encryption,omitempty"`
	ExtraProperties                  *map[string]interface{}           `json:"extraProperties,omitempty"`
	ManagedResourceGroupName         *string                           `json:"managedResourceGroupName,omitempty"`
	ManagedVirtualNetwork            *string                           `json:"managedVirtualNetwork,omitempty"`
	ManagedVirtualNetworkSettings    *ManagedVirtualNetworkSettings    `json:"managedVirtualNetworkSettings,omitempty"`
	PrivateEndpointConnections       *[]PrivateEndpointConnection      `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                *string                           `json:"provisioningState,omitempty"`
	PurviewConfiguration             *PurviewConfiguration             `json:"purviewConfiguration,omitempty"`
	SqlAdministratorLogin            *string                           `json:"sqlAdministratorLogin,omitempty"`
	SqlAdministratorLoginPassword    *string                           `json:"sqlAdministratorLoginPassword,omitempty"`
	VirtualNetworkProfile            *VirtualNetworkProfile            `json:"virtualNetworkProfile,omitempty"`
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfiguration `json:"workspaceRepositoryConfiguration,omitempty"`
	WorkspaceUID                     *string                           `json:"workspaceUID,omitempty"`
}
