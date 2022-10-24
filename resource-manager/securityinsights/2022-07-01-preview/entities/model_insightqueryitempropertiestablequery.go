package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightQueryItemPropertiesTableQuery struct {
	ColumnsDefinitions *[]InsightQueryItemPropertiesTableQueryColumnsDefinitionsInlined `json:"columnsDefinitions,omitempty"`
	QueriesDefinitions *[]InsightQueryItemPropertiesTableQueryQueriesDefinitionsInlined `json:"queriesDefinitions,omitempty"`
}
