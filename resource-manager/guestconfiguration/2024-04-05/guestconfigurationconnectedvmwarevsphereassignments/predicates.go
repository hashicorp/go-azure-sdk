package guestconfigurationconnectedvmwarevsphereassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GuestConfigurationAssignmentOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p GuestConfigurationAssignmentOperationPredicate) Matches(input GuestConfigurationAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && *p.Name != input.Name {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type GuestConfigurationAssignmentReportOperationPredicate struct {
	Id   *string
	Name *string
}

func (p GuestConfigurationAssignmentReportOperationPredicate) Matches(input GuestConfigurationAssignmentReport) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}
