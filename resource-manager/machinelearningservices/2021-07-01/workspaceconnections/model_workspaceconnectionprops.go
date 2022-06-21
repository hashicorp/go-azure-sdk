package workspaceconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionProps struct {
	AuthType    *string      `json:"authType,omitempty"`
	Category    *string      `json:"category,omitempty"`
	Target      *string      `json:"target,omitempty"`
	Value       *string      `json:"value,omitempty"`
	ValueFormat *ValueFormat `json:"valueFormat,omitempty"`
}
