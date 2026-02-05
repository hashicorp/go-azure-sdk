package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskDependencies struct {
	TaskIdRanges *[]TaskIdRange `json:"taskIdRanges,omitempty"`
	TaskIds      *[]string      `json:"taskIds,omitempty"`
}
