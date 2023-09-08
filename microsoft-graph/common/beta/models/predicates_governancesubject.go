package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceSubjectOperationPredicate struct {
	DisplayName   *string
	Email         *string
	Id            *string
	ODataType     *string
	PrincipalName *string
	Type          *string
}

func (p GovernanceSubjectOperationPredicate) Matches(input GovernanceSubject) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrincipalName != nil && (input.PrincipalName == nil || *p.PrincipalName != *input.PrincipalName) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
