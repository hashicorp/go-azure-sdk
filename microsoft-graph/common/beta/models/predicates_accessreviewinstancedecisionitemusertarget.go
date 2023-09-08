package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemUserTargetOperationPredicate struct {
	ODataType         *string
	UserDisplayName   *string
	UserId            *string
	UserPrincipalName *string
}

func (p AccessReviewInstanceDecisionItemUserTargetOperationPredicate) Matches(input AccessReviewInstanceDecisionItemUserTarget) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
