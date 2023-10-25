package distributedavailabilitygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateProperties struct {
	PrimaryAvailabilityGroupName   *string              `json:"primaryAvailabilityGroupName,omitempty"`
	ReplicationMode                *ReplicationModeType `json:"replicationMode,omitempty"`
	SecondaryAvailabilityGroupName *string              `json:"secondaryAvailabilityGroupName,omitempty"`
	SourceEndpoint                 *string              `json:"sourceEndpoint,omitempty"`
	TargetDatabase                 *string              `json:"targetDatabase,omitempty"`
}
