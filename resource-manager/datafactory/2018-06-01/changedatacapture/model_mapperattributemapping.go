package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MapperAttributeMapping struct {
	AttributeReference  *MapperAttributeReference   `json:"attributeReference,omitempty"`
	AttributeReferences *[]MapperAttributeReference `json:"attributeReferences,omitempty"`
	Expression          *string                     `json:"expression,omitempty"`
	FunctionName        *string                     `json:"functionName,omitempty"`
	Name                *string                     `json:"name,omitempty"`
	Type                *MappingType                `json:"type,omitempty"`
}
