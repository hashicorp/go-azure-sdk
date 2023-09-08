package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppRoleAssignmentOperationPredicate struct {
	AppRoleId            *string
	CreatedDateTime      *string
	DeletedDateTime      *string
	Id                   *string
	ODataType            *string
	PrincipalDisplayName *string
	PrincipalId          *string
	PrincipalType        *string
	ResourceDisplayName  *string
	ResourceId           *string
}

func (p AppRoleAssignmentOperationPredicate) Matches(input AppRoleAssignment) bool {

	if p.AppRoleId != nil && (input.AppRoleId == nil || *p.AppRoleId != *input.AppRoleId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrincipalDisplayName != nil && (input.PrincipalDisplayName == nil || *p.PrincipalDisplayName != *input.PrincipalDisplayName) {
		return false
	}

	if p.PrincipalId != nil && (input.PrincipalId == nil || *p.PrincipalId != *input.PrincipalId) {
		return false
	}

	if p.PrincipalType != nil && (input.PrincipalType == nil || *p.PrincipalType != *input.PrincipalType) {
		return false
	}

	if p.ResourceDisplayName != nil && (input.ResourceDisplayName == nil || *p.ResourceDisplayName != *input.ResourceDisplayName) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	return true
}
