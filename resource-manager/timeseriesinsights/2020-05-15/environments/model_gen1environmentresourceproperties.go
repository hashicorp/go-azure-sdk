package environments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Gen1EnvironmentResourceProperties struct {
	CreationTime                 *string                       `json:"creationTime,omitempty"`
	DataAccessFqdn               *string                       `json:"dataAccessFqdn,omitempty"`
	DataAccessId                 *string                       `json:"dataAccessId,omitempty"`
	DataRetentionTime            string                        `json:"dataRetentionTime"`
	PartitionKeyProperties       *[]TimeSeriesIdProperty       `json:"partitionKeyProperties,omitempty"`
	ProvisioningState            *ProvisioningState            `json:"provisioningState,omitempty"`
	Status                       *EnvironmentStatus            `json:"status,omitempty"`
	StorageLimitExceededBehavior *StorageLimitExceededBehavior `json:"storageLimitExceededBehavior,omitempty"`
}
