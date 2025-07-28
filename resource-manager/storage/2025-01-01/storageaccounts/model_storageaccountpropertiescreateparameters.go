package storageaccounts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountPropertiesCreateParameters struct {
	AccessTier                            *AccessTier                            `json:"accessTier,omitempty"`
	AllowBlobPublicAccess                 *bool                                  `json:"allowBlobPublicAccess,omitempty"`
	AllowCrossTenantReplication           *bool                                  `json:"allowCrossTenantReplication,omitempty"`
	AllowSharedKeyAccess                  *bool                                  `json:"allowSharedKeyAccess,omitempty"`
	AllowedCopyScope                      *AllowedCopyScope                      `json:"allowedCopyScope,omitempty"`
	AzureFilesIdentityBasedAuthentication *AzureFilesIdentityBasedAuthentication `json:"azureFilesIdentityBasedAuthentication,omitempty"`
	CustomDomain                          *CustomDomain                          `json:"customDomain,omitempty"`
	DefaultToOAuthAuthentication          *bool                                  `json:"defaultToOAuthAuthentication,omitempty"`
	DeletedAccountCreationTime            *string                                `json:"deletedAccountCreationTime,omitempty"`
	DnsEndpointType                       *DnsEndpointType                       `json:"dnsEndpointType,omitempty"`
	EnableExtendedGroups                  *bool                                  `json:"enableExtendedGroups,omitempty"`
	Encryption                            *Encryption                            `json:"encryption,omitempty"`
	ImmutableStorageWithVersioning        *ImmutableStorageAccount               `json:"immutableStorageWithVersioning,omitempty"`
	IsHnsEnabled                          *bool                                  `json:"isHnsEnabled,omitempty"`
	IsLocalUserEnabled                    *bool                                  `json:"isLocalUserEnabled,omitempty"`
	IsNfsV3Enabled                        *bool                                  `json:"isNfsV3Enabled,omitempty"`
	IsSftpEnabled                         *bool                                  `json:"isSftpEnabled,omitempty"`
	KeyPolicy                             *KeyPolicy                             `json:"keyPolicy,omitempty"`
	LargeFileSharesState                  *LargeFileSharesState                  `json:"largeFileSharesState,omitempty"`
	MinimumTlsVersion                     *MinimumTlsVersion                     `json:"minimumTlsVersion,omitempty"`
	NetworkAcls                           *NetworkRuleSet                        `json:"networkAcls,omitempty"`
	PublicNetworkAccess                   *PublicNetworkAccess                   `json:"publicNetworkAccess,omitempty"`
	RoutingPreference                     *RoutingPreference                     `json:"routingPreference,omitempty"`
	SasPolicy                             *SasPolicy                             `json:"sasPolicy,omitempty"`
	SupportsHTTPSTrafficOnly              *bool                                  `json:"supportsHttpsTrafficOnly,omitempty"`
}

func (o *StorageAccountPropertiesCreateParameters) GetDeletedAccountCreationTimeAsTime() (*time.Time, error) {
	if o.DeletedAccountCreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeletedAccountCreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *StorageAccountPropertiesCreateParameters) SetDeletedAccountCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeletedAccountCreationTime = &formatted
}
