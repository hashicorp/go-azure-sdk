package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProgramControlTypeOperationPredicate struct {
	ControlTypeGroupId *string
	DisplayName        *string
	Id                 *string
	ODataType          *string
}

func (p ProgramControlTypeOperationPredicate) Matches(input ProgramControlType) bool {

	if p.ControlTypeGroupId != nil && (input.ControlTypeGroupId == nil || *p.ControlTypeGroupId != *input.ControlTypeGroupId) {
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

	return true
}
