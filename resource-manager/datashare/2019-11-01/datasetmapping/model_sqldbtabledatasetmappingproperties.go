package datasetmapping

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlDBTableDataSetMappingProperties struct {
	DataSetId            string                `json:"dataSetId"`
	DataSetMappingStatus *DataSetMappingStatus `json:"dataSetMappingStatus,omitempty"`
	DatabaseName         string                `json:"databaseName"`
	ProvisioningState    *ProvisioningState    `json:"provisioningState,omitempty"`
	SchemaName           string                `json:"schemaName"`
	SqlServerResourceId  string                `json:"sqlServerResourceId"`
	TableName            string                `json:"tableName"`
}
