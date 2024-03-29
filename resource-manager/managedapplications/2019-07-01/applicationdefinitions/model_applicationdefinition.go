package applicationdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationDefinition struct {
	Id         *string                         `json:"id,omitempty"`
	Location   *string                         `json:"location,omitempty"`
	ManagedBy  *string                         `json:"managedBy,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties ApplicationDefinitionProperties `json:"properties"`
	Sku        *Sku                            `json:"sku,omitempty"`
	Tags       *map[string]string              `json:"tags,omitempty"`
	Type       *string                         `json:"type,omitempty"`
}
