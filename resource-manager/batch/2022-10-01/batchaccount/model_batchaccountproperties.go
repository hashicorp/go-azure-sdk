package batchaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchAccountProperties struct {
	AccountEndpoint                       *string                          `json:"accountEndpoint,omitempty"`
	ActiveJobAndJobScheduleQuota          *int64                           `json:"activeJobAndJobScheduleQuota,omitempty"`
	AllowedAuthenticationModes            *[]AuthenticationMode            `json:"allowedAuthenticationModes,omitempty"`
	AutoStorage                           *AutoStorageProperties           `json:"autoStorage"`
	DedicatedCoreQuota                    *int64                           `json:"dedicatedCoreQuota,omitempty"`
	DedicatedCoreQuotaPerVMFamily         *[]VirtualMachineFamilyCoreQuota `json:"dedicatedCoreQuotaPerVMFamily,omitempty"`
	DedicatedCoreQuotaPerVMFamilyEnforced *bool                            `json:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty"`
	Encryption                            *EncryptionProperties            `json:"encryption"`
	KeyVaultReference                     *KeyVaultReference               `json:"keyVaultReference"`
	LowPriorityCoreQuota                  *int64                           `json:"lowPriorityCoreQuota,omitempty"`
	NetworkProfile                        *NetworkProfile                  `json:"networkProfile"`
	NodeManagementEndpoint                *string                          `json:"nodeManagementEndpoint,omitempty"`
	PoolAllocationMode                    *PoolAllocationMode              `json:"poolAllocationMode,omitempty"`
	PoolQuota                             *int64                           `json:"poolQuota,omitempty"`
	PrivateEndpointConnections            *[]PrivateEndpointConnection     `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                     *ProvisioningState               `json:"provisioningState,omitempty"`
	PublicNetworkAccess                   *PublicNetworkAccessType         `json:"publicNetworkAccess,omitempty"`
}
