package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountPropertiesCreateParameters struct {
	AccessTier                            *AccessTier                            `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                  `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                  `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                  `json:"allowSharedKeyAccess,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	CustomDomain                          *CustomDomain                          `json:"customDomain,omitempty"`
	Encryption                            *Encryption                            `json:"encryption,omitempty"`
	IsHnsEnabled                          *bool                                  `json:"isHnsEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                  `json:"isNfsV3Enabled,omitempty"`
	KeyPolicy                             *KeyPolicy                             `json:"keyPolicy,omitempty"`
	LargeFileSharesState                  *LargeFileSharesState                  `json:"largeFileSharesState,omitempty"`
	MinimumTlsVersion                     *MinimumTlsVersion                     `json:"minimumTlsVersion,omitempty"`
	NetworkAcls                           *NetworkRuleSet                        `json:"networkAcls,omitempty"`
	RoutingPreference                     *RoutingPreference                     `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy                             `json:"sasPolicy,omitempty"`
	SupportsHTTPSTrafficOnly              *bool                                  `json:"supportsHttpsTrafficOnly,omitempty"`
}
