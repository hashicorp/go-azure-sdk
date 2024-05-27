package dataproducts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductProperties struct {
	AvailableMinorVersions              *[]string                          `json:"availableMinorVersions,omitempty"`
	ConsumptionEndpoints                *ConsumptionEndpointsProperties    `json:"consumptionEndpoints,omitempty"`
	CurrentMinorVersion                 *string                            `json:"currentMinorVersion,omitempty"`
	CustomerEncryptionKey               *EncryptionKeyDetails              `json:"customerEncryptionKey,omitempty"`
	CustomerManagedKeyEncryptionEnabled *ControlState                      `json:"customerManagedKeyEncryptionEnabled,omitempty"`
	Documentation                       *string                            `json:"documentation,omitempty"`
	KeyVaultUrl                         *string                            `json:"keyVaultUrl,omitempty"`
	MajorVersion                        string                             `json:"majorVersion"`
	ManagedResourceGroupConfiguration   *ManagedResourceGroupConfiguration `json:"managedResourceGroupConfiguration,omitempty"`
	Networkacls                         *DataProductNetworkAcls            `json:"networkacls,omitempty"`
	Owners                              *[]string                          `json:"owners,omitempty"`
	PrivateLinksEnabled                 *ControlState                      `json:"privateLinksEnabled,omitempty"`
	Product                             string                             `json:"product"`
	ProvisioningState                   *ProvisioningState                 `json:"provisioningState,omitempty"`
	PublicNetworkAccess                 *ControlState                      `json:"publicNetworkAccess,omitempty"`
	Publisher                           string                             `json:"publisher"`
	PurviewAccount                      *string                            `json:"purviewAccount,omitempty"`
	PurviewCollection                   *string                            `json:"purviewCollection,omitempty"`
	Redundancy                          *ControlState                      `json:"redundancy,omitempty"`
	ResourceGuid                        *string                            `json:"resourceGuid,omitempty"`
}
