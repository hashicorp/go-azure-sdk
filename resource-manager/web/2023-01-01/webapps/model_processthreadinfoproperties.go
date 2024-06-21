package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessThreadInfoProperties struct {
	BasePriority       *int64  `json:"base_priority,omitempty"`
	CurrentPriority    *int64  `json:"current_priority,omitempty"`
	Href               *string `json:"href,omitempty"`
	Identifier         *int64  `json:"identifier,omitempty"`
	PriorityLevel      *string `json:"priority_level,omitempty"`
	Process            *string `json:"process,omitempty"`
	StartAddress       *string `json:"start_address,omitempty"`
	StartTime          *string `json:"start_time,omitempty"`
	State              *string `json:"state,omitempty"`
	TotalProcessorTime *string `json:"total_processor_time,omitempty"`
	UserProcessorTime  *string `json:"user_processor_time,omitempty"`
	WaitReason         *string `json:"wait_reason,omitempty"`
}
