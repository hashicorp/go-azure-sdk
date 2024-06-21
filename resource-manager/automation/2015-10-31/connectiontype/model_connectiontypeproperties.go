package connectiontype

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionTypeProperties struct {
	CreationTime     *string                     `json:"creationTime,omitempty"`
	Description      *string                     `json:"description,omitempty"`
	FieldDefinitions *map[string]FieldDefinition `json:"fieldDefinitions,omitempty"`
	IsGlobal         *bool                       `json:"isGlobal,omitempty"`
	LastModifiedTime *string                     `json:"lastModifiedTime,omitempty"`
}
