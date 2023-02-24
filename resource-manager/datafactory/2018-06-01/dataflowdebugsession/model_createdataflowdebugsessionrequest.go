package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDataFlowDebugSessionRequest struct {
	ComputeType        *string                          `json:"computeType,omitempty"`
	CoreCount          *int64                           `json:"coreCount,omitempty"`
	IntegrationRuntime *IntegrationRuntimeDebugResource `json:"integrationRuntime,omitempty"`
	TimeToLive         *int64                           `json:"timeToLive,omitempty"`
}
