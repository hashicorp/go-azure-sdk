package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CassandraTableDatasetTypeProperties struct {
	Keyspace  *interface{} `json:"keyspace,omitempty"`
	TableName *interface{} `json:"tableName,omitempty"`
}
