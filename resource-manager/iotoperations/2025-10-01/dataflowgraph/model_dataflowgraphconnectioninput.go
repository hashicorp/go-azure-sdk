package dataflowgraph

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphConnectionInput struct {
	Name   string                                 `json:"name"`
	Schema *DataflowGraphConnectionSchemaSettings `json:"schema,omitempty"`
}
