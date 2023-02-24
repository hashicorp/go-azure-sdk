package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugSessionInfo struct {
	ComputeType            *string `json:"computeType,omitempty"`
	CoreCount              *int64  `json:"coreCount,omitempty"`
	DataFlowName           *string `json:"dataFlowName,omitempty"`
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty"`
	LastActivityTime       *string `json:"lastActivityTime,omitempty"`
	NodeCount              *int64  `json:"nodeCount,omitempty"`
	SessionId              *string `json:"sessionId,omitempty"`
	StartTime              *string `json:"startTime,omitempty"`
	TimeToLiveInMinutes    *int64  `json:"timeToLiveInMinutes,omitempty"`
}
