package syncgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncFullSchemaTableColumn struct {
	DataSize     *string `json:"dataSize,omitempty"`
	DataType     *string `json:"dataType,omitempty"`
	ErrorId      *string `json:"errorId,omitempty"`
	HasError     *bool   `json:"hasError,omitempty"`
	IsPrimaryKey *bool   `json:"isPrimaryKey,omitempty"`
	Name         *string `json:"name,omitempty"`
	QuotedName   *string `json:"quotedName,omitempty"`
}
