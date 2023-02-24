package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugCommandPayload struct {
	Columns    *[]string `json:"columns,omitempty"`
	Expression *string   `json:"expression,omitempty"`
	RowLimits  *int64    `json:"rowLimits,omitempty"`
	StreamName string    `json:"streamName"`
}
