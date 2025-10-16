package tuningoptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectRecommendationDetails struct {
	DatabaseName    *string   `json:"databaseName,omitempty"`
	IncludedColumns *[]string `json:"includedColumns,omitempty"`
	IndexColumns    *[]string `json:"indexColumns,omitempty"`
	IndexName       *string   `json:"indexName,omitempty"`
	IndexType       *string   `json:"indexType,omitempty"`
	Schema          *string   `json:"schema,omitempty"`
	Table           *string   `json:"table,omitempty"`
}
