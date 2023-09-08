package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationCustomizationOperationPredicate struct {
	AllowDisplayNameUpdate   *bool
	IsSyncDeferred           *bool
	ODataType                *string
	SynchronizationStartDate *string
}

func (p EducationSynchronizationCustomizationOperationPredicate) Matches(input EducationSynchronizationCustomization) bool {

	if p.AllowDisplayNameUpdate != nil && (input.AllowDisplayNameUpdate == nil || *p.AllowDisplayNameUpdate != *input.AllowDisplayNameUpdate) {
		return false
	}

	if p.IsSyncDeferred != nil && (input.IsSyncDeferred == nil || *p.IsSyncDeferred != *input.IsSyncDeferred) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SynchronizationStartDate != nil && (input.SynchronizationStartDate == nil || *p.SynchronizationStartDate != *input.SynchronizationStartDate) {
		return false
	}

	return true
}
