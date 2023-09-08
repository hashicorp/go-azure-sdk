package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleOperationPredicate struct {
	DeletedDateTime *string
	Description     *string
	DisplayName     *string
	Id              *string
	ODataType       *string
	RoleTemplateId  *string
}

func (p DirectoryRoleOperationPredicate) Matches(input DirectoryRole) bool {

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RoleTemplateId != nil && (input.RoleTemplateId == nil || *p.RoleTemplateId != *input.RoleTemplateId) {
		return false
	}

	return true
}
