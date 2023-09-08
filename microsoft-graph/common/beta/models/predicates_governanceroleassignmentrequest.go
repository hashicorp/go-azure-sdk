package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleAssignmentRequestOperationPredicate struct {
	AssignmentState                *string
	Id                             *string
	LinkedEligibleRoleAssignmentId *string
	ODataType                      *string
	Reason                         *string
	RequestedDateTime              *string
	ResourceId                     *string
	RoleDefinitionId               *string
	SubjectId                      *string
	Type                           *string
}

func (p GovernanceRoleAssignmentRequestOperationPredicate) Matches(input GovernanceRoleAssignmentRequest) bool {

	if p.AssignmentState != nil && (input.AssignmentState == nil || *p.AssignmentState != *input.AssignmentState) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LinkedEligibleRoleAssignmentId != nil && (input.LinkedEligibleRoleAssignmentId == nil || *p.LinkedEligibleRoleAssignmentId != *input.LinkedEligibleRoleAssignmentId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Reason != nil && (input.Reason == nil || *p.Reason != *input.Reason) {
		return false
	}

	if p.RequestedDateTime != nil && (input.RequestedDateTime == nil || *p.RequestedDateTime != *input.RequestedDateTime) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	if p.RoleDefinitionId != nil && (input.RoleDefinitionId == nil || *p.RoleDefinitionId != *input.RoleDefinitionId) {
		return false
	}

	if p.SubjectId != nil && (input.SubjectId == nil || *p.SubjectId != *input.SubjectId) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
