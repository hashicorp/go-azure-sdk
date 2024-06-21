package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceProperties struct {
	Authorizations             *[]WorkspaceProviderAuthorization `json:"authorizations,omitempty"`
	CreatedBy                  *CreatedBy                        `json:"createdBy,omitempty"`
	CreatedDateTime            *string                           `json:"createdDateTime,omitempty"`
	DiskEncryptionSetId        *string                           `json:"diskEncryptionSetId,omitempty"`
	Encryption                 *WorkspacePropertiesEncryption    `json:"encryption,omitempty"`
	ManagedDiskIdentity        *ManagedIdentityConfiguration     `json:"managedDiskIdentity,omitempty"`
	ManagedResourceGroupId     string                            `json:"managedResourceGroupId"`
	Parameters                 *WorkspaceCustomParameters        `json:"parameters,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection      `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState                `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess              `json:"publicNetworkAccess,omitempty"`
	RequiredNsgRules           *RequiredNsgRules                 `json:"requiredNsgRules,omitempty"`
	StorageAccountIdentity     *ManagedIdentityConfiguration     `json:"storageAccountIdentity,omitempty"`
	UiDefinitionUri            *string                           `json:"uiDefinitionUri,omitempty"`
	UpdatedBy                  *CreatedBy                        `json:"updatedBy,omitempty"`
	WorkspaceId                *string                           `json:"workspaceId,omitempty"`
	WorkspaceUrl               *string                           `json:"workspaceUrl,omitempty"`
}
