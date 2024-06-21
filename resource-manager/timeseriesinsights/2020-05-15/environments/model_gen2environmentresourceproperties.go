package environments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Gen2EnvironmentResourceProperties struct {
	CreationTime           *string                           `json:"creationTime,omitempty"`
	DataAccessFqdn         *string                           `json:"dataAccessFqdn,omitempty"`
	DataAccessId           *string                           `json:"dataAccessId,omitempty"`
	ProvisioningState      *ProvisioningState                `json:"provisioningState,omitempty"`
	Status                 *EnvironmentStatus                `json:"status,omitempty"`
	StorageConfiguration   Gen2StorageConfigurationOutput    `json:"storageConfiguration"`
	TimeSeriesIdProperties []TimeSeriesIdProperty            `json:"timeSeriesIdProperties"`
	WarmStoreConfiguration *WarmStoreConfigurationProperties `json:"warmStoreConfiguration,omitempty"`
}
