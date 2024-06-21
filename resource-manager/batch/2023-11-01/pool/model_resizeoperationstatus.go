package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResizeOperationStatus struct {
	Errors                 *[]ResizeError                 `json:"errors,omitempty"`
	NodeDeallocationOption *ComputeNodeDeallocationOption `json:"nodeDeallocationOption,omitempty"`
	ResizeTimeout          *string                        `json:"resizeTimeout,omitempty"`
	StartTime              *string                        `json:"startTime,omitempty"`
	TargetDedicatedNodes   *int64                         `json:"targetDedicatedNodes,omitempty"`
	TargetLowPriorityNodes *int64                         `json:"targetLowPriorityNodes,omitempty"`
}
