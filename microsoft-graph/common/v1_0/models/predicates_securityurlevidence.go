package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidenceOperationPredicate struct {
	CreatedDateTime          *string
	ODataType                *string
	RemediationStatusDetails *string
	Url                      *string
}

func (p SecurityUrlEvidenceOperationPredicate) Matches(input SecurityUrlEvidence) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
