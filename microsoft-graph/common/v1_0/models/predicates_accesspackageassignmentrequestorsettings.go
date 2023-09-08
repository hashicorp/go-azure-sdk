package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestorSettingsOperationPredicate struct {
	AllowCustomAssignmentSchedule          *bool
	EnableOnBehalfRequestorsToAddAccess    *bool
	EnableOnBehalfRequestorsToRemoveAccess *bool
	EnableOnBehalfRequestorsToUpdateAccess *bool
	EnableTargetsToSelfAddAccess           *bool
	EnableTargetsToSelfRemoveAccess        *bool
	EnableTargetsToSelfUpdateAccess        *bool
	ODataType                              *string
}

func (p AccessPackageAssignmentRequestorSettingsOperationPredicate) Matches(input AccessPackageAssignmentRequestorSettings) bool {

	if p.AllowCustomAssignmentSchedule != nil && (input.AllowCustomAssignmentSchedule == nil || *p.AllowCustomAssignmentSchedule != *input.AllowCustomAssignmentSchedule) {
		return false
	}

	if p.EnableOnBehalfRequestorsToAddAccess != nil && (input.EnableOnBehalfRequestorsToAddAccess == nil || *p.EnableOnBehalfRequestorsToAddAccess != *input.EnableOnBehalfRequestorsToAddAccess) {
		return false
	}

	if p.EnableOnBehalfRequestorsToRemoveAccess != nil && (input.EnableOnBehalfRequestorsToRemoveAccess == nil || *p.EnableOnBehalfRequestorsToRemoveAccess != *input.EnableOnBehalfRequestorsToRemoveAccess) {
		return false
	}

	if p.EnableOnBehalfRequestorsToUpdateAccess != nil && (input.EnableOnBehalfRequestorsToUpdateAccess == nil || *p.EnableOnBehalfRequestorsToUpdateAccess != *input.EnableOnBehalfRequestorsToUpdateAccess) {
		return false
	}

	if p.EnableTargetsToSelfAddAccess != nil && (input.EnableTargetsToSelfAddAccess == nil || *p.EnableTargetsToSelfAddAccess != *input.EnableTargetsToSelfAddAccess) {
		return false
	}

	if p.EnableTargetsToSelfRemoveAccess != nil && (input.EnableTargetsToSelfRemoveAccess == nil || *p.EnableTargetsToSelfRemoveAccess != *input.EnableTargetsToSelfRemoveAccess) {
		return false
	}

	if p.EnableTargetsToSelfUpdateAccess != nil && (input.EnableTargetsToSelfUpdateAccess == nil || *p.EnableTargetsToSelfUpdateAccess != *input.EnableTargetsToSelfUpdateAccess) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
