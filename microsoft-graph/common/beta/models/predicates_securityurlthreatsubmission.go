package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionOperationPredicate struct {
	CreatedDateTime *string
	Id              *string
	ODataType       *string
	TenantId        *string
	WebUrl          *string
}

func (p SecurityUrlThreatSubmissionOperationPredicate) Matches(input SecurityUrlThreatSubmission) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
