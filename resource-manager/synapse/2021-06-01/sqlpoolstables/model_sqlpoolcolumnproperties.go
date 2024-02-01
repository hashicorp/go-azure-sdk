package sqlpoolstables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolColumnProperties struct {
	ColumnType *ColumnDataType `json:"columnType,omitempty"`
	IsComputed *bool           `json:"isComputed,omitempty"`
}
