package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GoogleBigQueryDatasetTypeProperties struct {
	Dataset   *interface{} `json:"dataset,omitempty"`
	Table     *interface{} `json:"table,omitempty"`
	TableName *interface{} `json:"tableName,omitempty"`
}
