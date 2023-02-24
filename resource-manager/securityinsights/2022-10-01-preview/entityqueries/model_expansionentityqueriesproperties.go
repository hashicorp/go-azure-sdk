package entityqueries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpansionEntityQueriesProperties struct {
	DataSources       *[]string     `json:"dataSources,omitempty"`
	DisplayName       *string       `json:"displayName,omitempty"`
	InputEntityType   *EntityType   `json:"inputEntityType,omitempty"`
	InputFields       *[]string     `json:"inputFields,omitempty"`
	OutputEntityTypes *[]EntityType `json:"outputEntityTypes,omitempty"`
	QueryTemplate     *string       `json:"queryTemplate,omitempty"`
}
