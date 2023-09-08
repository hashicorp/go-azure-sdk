package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeDefinitionOperationPredicate struct {
	Anchor         *bool
	CaseExact      *bool
	DefaultValue   *string
	FlowNullValues *bool
	Multivalued    *bool
	Name           *string
	ODataType      *string
	Required       *bool
}

func (p AttributeDefinitionOperationPredicate) Matches(input AttributeDefinition) bool {

	if p.Anchor != nil && (input.Anchor == nil || *p.Anchor != *input.Anchor) {
		return false
	}

	if p.CaseExact != nil && (input.CaseExact == nil || *p.CaseExact != *input.CaseExact) {
		return false
	}

	if p.DefaultValue != nil && (input.DefaultValue == nil || *p.DefaultValue != *input.DefaultValue) {
		return false
	}

	if p.FlowNullValues != nil && (input.FlowNullValues == nil || *p.FlowNullValues != *input.FlowNullValues) {
		return false
	}

	if p.Multivalued != nil && (input.Multivalued == nil || *p.Multivalued != *input.Multivalued) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Required != nil && (input.Required == nil || *p.Required != *input.Required) {
		return false
	}

	return true
}
