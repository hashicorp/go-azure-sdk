package alertruletemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityMapping struct {
	EntityType    *EntityMappingType `json:"entityType,omitempty"`
	FieldMappings *[]FieldMapping    `json:"fieldMappings,omitempty"`
}
