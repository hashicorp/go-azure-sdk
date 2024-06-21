package automationaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationAccountProperties struct {
	CreationTime               *string                      `json:"creationTime,omitempty"`
	Description                *string                      `json:"description,omitempty"`
	Encryption                 *EncryptionProperties        `json:"encryption,omitempty"`
	LastModifiedBy             *string                      `json:"lastModifiedBy,omitempty"`
	LastModifiedTime           *string                      `json:"lastModifiedTime,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	PublicNetworkAccess        *bool                        `json:"publicNetworkAccess,omitempty"`
	Sku                        *Sku                         `json:"sku,omitempty"`
	State                      *AutomationAccountState      `json:"state,omitempty"`
}
