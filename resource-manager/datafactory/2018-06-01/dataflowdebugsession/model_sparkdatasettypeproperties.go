package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkDatasetTypeProperties struct {
	Schema    *string      `json:"schema,omitempty"`
	Table     *string      `json:"table,omitempty"`
	TableName *interface{} `json:"tableName,omitempty"`
}
