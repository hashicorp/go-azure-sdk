package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugCommandResponse struct {
	Data   *string `json:"data,omitempty"`
	Status *string `json:"status,omitempty"`
}
