package configurationstores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationStoreProperties struct {
	CreateMode                 *CreateMode                           `json:"createMode,omitempty"`
	CreationDate               *string                               `json:"creationDate,omitempty"`
	DisableLocalAuth           *bool                                 `json:"disableLocalAuth,omitempty"`
	EnablePurgeProtection      *bool                                 `json:"enablePurgeProtection,omitempty"`
	Encryption                 *EncryptionProperties                 `json:"encryption,omitempty"`
	Endpoint                   *string                               `json:"endpoint,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnectionReference `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState                    `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess                  `json:"publicNetworkAccess,omitempty"`
	SoftDeleteRetentionInDays  *int64                                `json:"softDeleteRetentionInDays,omitempty"`
}
