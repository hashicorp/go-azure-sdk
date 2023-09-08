package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemServicePrincipalTargetOperationPredicate struct {
	AppId                       *string
	ODataType                   *string
	ServicePrincipalDisplayName *string
	ServicePrincipalId          *string
}

func (p AccessReviewInstanceDecisionItemServicePrincipalTargetOperationPredicate) Matches(input AccessReviewInstanceDecisionItemServicePrincipalTarget) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServicePrincipalDisplayName != nil && (input.ServicePrincipalDisplayName == nil || *p.ServicePrincipalDisplayName != *input.ServicePrincipalDisplayName) {
		return false
	}

	if p.ServicePrincipalId != nil && (input.ServicePrincipalId == nil || *p.ServicePrincipalId != *input.ServicePrincipalId) {
		return false
	}

	return true
}
