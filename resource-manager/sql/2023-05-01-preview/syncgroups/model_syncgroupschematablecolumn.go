package syncgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupSchemaTableColumn struct {
	DataSize   *string `json:"dataSize,omitempty"`
	DataType   *string `json:"dataType,omitempty"`
	QuotedName *string `json:"quotedName,omitempty"`
}
