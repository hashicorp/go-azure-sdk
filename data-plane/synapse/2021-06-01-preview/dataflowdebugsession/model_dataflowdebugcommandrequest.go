package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugCommandRequest struct {
	Command        *DataFlowDebugCommandType    `json:"command,omitempty"`
	CommandPayload *DataFlowDebugCommandPayload `json:"commandPayload,omitempty"`
	SessionId      *string                      `json:"sessionId,omitempty"`
}
