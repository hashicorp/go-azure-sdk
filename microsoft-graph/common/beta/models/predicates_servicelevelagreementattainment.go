package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceLevelAgreementAttainmentOperationPredicate struct {
	EndDate   *string
	ODataType *string
	StartDate *string
}

func (p ServiceLevelAgreementAttainmentOperationPredicate) Matches(input ServiceLevelAgreementAttainment) bool {

	if p.EndDate != nil && (input.EndDate == nil || *p.EndDate != *input.EndDate) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDate != nil && (input.StartDate == nil || *p.StartDate != *input.StartDate) {
		return false
	}

	return true
}
