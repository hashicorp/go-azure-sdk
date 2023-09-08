package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationUserOperationPredicate struct {
	DisplayName     *string
	Email           *string
	ODataType       *string
	OutOfOfficeDays *int64
	UserId          *string
}

func (p AttackSimulationUserOperationPredicate) Matches(input AttackSimulationUser) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OutOfOfficeDays != nil && (input.OutOfOfficeDays == nil || *p.OutOfOfficeDays != *input.OutOfOfficeDays) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
