package entityqueries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityEntityQueryTemplateProperties struct {
	Content                 *string                                                `json:"content,omitempty"`
	DataTypes               *[]DataTypeDefinitions                                 `json:"dataTypes,omitempty"`
	Description             *string                                                `json:"description,omitempty"`
	EntitiesFilter          *map[string][]string                                   `json:"entitiesFilter,omitempty"`
	InputEntityType         *EntityType                                            `json:"inputEntityType,omitempty"`
	QueryDefinitions        *ActivityEntityQueryTemplatePropertiesQueryDefinitions `json:"queryDefinitions"`
	RequiredInputFieldsSets *[][]string                                            `json:"requiredInputFieldsSets,omitempty"`
	Title                   *string                                                `json:"title,omitempty"`
}
