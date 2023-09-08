package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleDefinitionOperationPredicate struct {
	Description *string
	DisplayName *string
	Id          *string
	IsBuiltIn   *bool
	IsEnabled   *bool
	ODataType   *string
	TemplateId  *string
	Version     *string
}

func (p UnifiedRoleDefinitionOperationPredicate) Matches(input UnifiedRoleDefinition) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsBuiltIn != nil && (input.IsBuiltIn == nil || *p.IsBuiltIn != *input.IsBuiltIn) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TemplateId != nil && (input.TemplateId == nil || *p.TemplateId != *input.TemplateId) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
