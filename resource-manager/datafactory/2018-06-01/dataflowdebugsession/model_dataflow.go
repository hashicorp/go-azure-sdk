package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlow struct {
	Annotations *[]interface{}  `json:"annotations,omitempty"`
	Description *string         `json:"description,omitempty"`
	Folder      *DataFlowFolder `json:"folder,omitempty"`
	Type        string          `json:"type"`
}
