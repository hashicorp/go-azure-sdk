package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsightQueryItemProperties struct {
	AdditionalQuery         *InsightQueryItemPropertiesAdditionalQuery    `json:"additionalQuery"`
	BaseQuery               *string                                       `json:"baseQuery,omitempty"`
	ChartQuery              *interface{}                                  `json:"chartQuery,omitempty"`
	DataTypes               *[]EntityQueryItemPropertiesDataTypesInlined  `json:"dataTypes,omitempty"`
	DefaultTimeRange        *InsightQueryItemPropertiesDefaultTimeRange   `json:"defaultTimeRange"`
	Description             *string                                       `json:"description,omitempty"`
	DisplayName             *string                                       `json:"displayName,omitempty"`
	EntitiesFilter          *interface{}                                  `json:"entitiesFilter,omitempty"`
	InputEntityType         *EntityType                                   `json:"inputEntityType,omitempty"`
	ReferenceTimeRange      *InsightQueryItemPropertiesReferenceTimeRange `json:"referenceTimeRange"`
	RequiredInputFieldsSets *[][]string                                   `json:"requiredInputFieldsSets,omitempty"`
	TableQuery              *InsightQueryItemPropertiesTableQuery         `json:"tableQuery"`
}
