package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinition struct {
	Anchor            *bool                               `json:"anchor,omitempty"`
	ApiExpressions    *[]StringKeyStringValuePair         `json:"apiExpressions,omitempty"`
	CaseExact         *bool                               `json:"caseExact,omitempty"`
	DefaultValue      *string                             `json:"defaultValue,omitempty"`
	FlowNullValues    *bool                               `json:"flowNullValues,omitempty"`
	Metadata          *[]AttributeDefinitionMetadataEntry `json:"metadata,omitempty"`
	Multivalued       *bool                               `json:"multivalued,omitempty"`
	Mutability        *AttributeDefinitionMutability      `json:"mutability,omitempty"`
	Name              *string                             `json:"name,omitempty"`
	ODataType         *string                             `json:"@odata.type,omitempty"`
	ReferencedObjects *[]ReferencedObject                 `json:"referencedObjects,omitempty"`
	Required          *bool                               `json:"required,omitempty"`
	Type              *AttributeDefinitionType            `json:"type,omitempty"`
}
