package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmServiceTypeHealthPolicy struct {
	MaxPercentUnhealthyPartitionsPerService *int64 `json:"maxPercentUnhealthyPartitionsPerService,omitempty"`
	MaxPercentUnhealthyReplicasPerPartition *int64 `json:"maxPercentUnhealthyReplicasPerPartition,omitempty"`
	MaxPercentUnhealthyServices             *int64 `json:"maxPercentUnhealthyServices,omitempty"`
}
