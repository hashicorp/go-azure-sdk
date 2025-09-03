package agents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentUpdateProperties struct {
	Description         *string              `json:"description,omitempty"`
	UploadLimitSchedule *UploadLimitSchedule `json:"uploadLimitSchedule,omitempty"`
}
