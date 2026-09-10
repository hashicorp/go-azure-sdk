package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataTableResponseObject struct {
	Columns   *[]DataTableResponseColumn `json:"columns,omitempty"`
	Rows      *[][]string                `json:"rows,omitempty"`
	TableName *string                    `json:"tableName,omitempty"`
}
