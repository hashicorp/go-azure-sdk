package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryDataFlowDebugSessionsResponse struct {
	NextLink *string                     `json:"nextLink,omitempty"`
	Value    *[]DataFlowDebugSessionInfo `json:"value,omitempty"`
}
