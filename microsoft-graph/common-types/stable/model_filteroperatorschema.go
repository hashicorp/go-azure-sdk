package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterOperatorSchema struct {
	Arity                     *FilterOperatorSchemaArity                     `json:"arity,omitempty"`
	Id                        *string                                        `json:"id,omitempty"`
	MultivaluedComparisonType *FilterOperatorSchemaMultivaluedComparisonType `json:"multivaluedComparisonType,omitempty"`
	ODataType                 *string                                        `json:"@odata.type,omitempty"`
	SupportedAttributeTypes   *FilterOperatorSchemaSupportedAttributeTypes   `json:"supportedAttributeTypes,omitempty"`
}
