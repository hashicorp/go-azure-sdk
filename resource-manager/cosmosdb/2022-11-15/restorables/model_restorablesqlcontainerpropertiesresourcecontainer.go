package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableSqlContainerPropertiesResourceContainer struct {
	AnalyticalStorageTtl     *int64                    `json:"analyticalStorageTtl,omitempty"`
	ClientEncryptionPolicy   *ClientEncryptionPolicy   `json:"clientEncryptionPolicy,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int64                    `json:"defaultTtl,omitempty"`
	Etag                     *string                   `json:"_etag,omitempty"`
	Id                       *string                   `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey    `json:"partitionKey,omitempty"`
	Rid                      *string                   `json:"_rid,omitempty"`
	Self                     *string                   `json:"_self,omitempty"`
	Ts                       *float64                  `json:"_ts,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `json:"uniqueKeyPolicy,omitempty"`
}
