package distributedavailabilitygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributedAvailabilityGroupProperties struct {
	DistributedAvailabilityGroupId *string          `json:"distributedAvailabilityGroupId,omitempty"`
	LastHardenedLsn                *string          `json:"lastHardenedLsn,omitempty"`
	LinkState                      *string          `json:"linkState,omitempty"`
	PrimaryAvailabilityGroupName   *string          `json:"primaryAvailabilityGroupName,omitempty"`
	ReplicationMode                *ReplicationMode `json:"replicationMode,omitempty"`
	SecondaryAvailabilityGroupName *string          `json:"secondaryAvailabilityGroupName,omitempty"`
	SourceEndpoint                 *string          `json:"sourceEndpoint,omitempty"`
	SourceReplicaId                *string          `json:"sourceReplicaId,omitempty"`
	TargetDatabase                 *string          `json:"targetDatabase,omitempty"`
	TargetReplicaId                *string          `json:"targetReplicaId,omitempty"`
}
