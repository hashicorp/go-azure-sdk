package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTypeHealthPolicy struct {
	MaxPercentUnhealthyPartitionsPerService int64 `json:"maxPercentUnhealthyPartitionsPerService"`
	MaxPercentUnhealthyReplicasPerPartition int64 `json:"maxPercentUnhealthyReplicasPerPartition"`
	MaxPercentUnhealthyServices             int64 `json:"maxPercentUnhealthyServices"`
}
