package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobUpdateParameter struct {
	Constraints        *JobConstraints     `json:"constraints,omitempty"`
	Metadata           *[]MetadataItem     `json:"metadata,omitempty"`
	OnAllTasksComplete *OnAllTasksComplete `json:"onAllTasksComplete,omitempty"`
	PoolInfo           PoolInformation     `json:"poolInfo"`
	Priority           *int64              `json:"priority,omitempty"`
}
