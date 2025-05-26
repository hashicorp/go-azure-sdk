package replicationprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVReplicaAzureManagedDiskDetails struct {
	DiskEncryptionSetId   *string          `json:"diskEncryptionSetId,omitempty"`
	DiskId                *string          `json:"diskId,omitempty"`
	DiskSizeInGB          *int64           `json:"diskSizeInGB,omitempty"`
	Iops                  *int64           `json:"iops,omitempty"`
	ReplicaDiskType       *string          `json:"replicaDiskType,omitempty"`
	SectorSizeInBytes     *int64           `json:"sectorSizeInBytes,omitempty"`
	SeedManagedDiskId     *string          `json:"seedManagedDiskId,omitempty"`
	TargetDiskAccountType *DiskAccountType `json:"targetDiskAccountType,omitempty"`
	ThroughputInMbps      *int64           `json:"throughputInMbps,omitempty"`
}
