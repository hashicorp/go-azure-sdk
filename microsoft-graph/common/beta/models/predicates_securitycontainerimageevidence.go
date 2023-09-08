package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceOperationPredicate struct {
	CreatedDateTime          *string
	ImageId                  *string
	ODataType                *string
	RemediationStatusDetails *string
}

func (p SecurityContainerImageEvidenceOperationPredicate) Matches(input SecurityContainerImageEvidence) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ImageId != nil && (input.ImageId == nil || *p.ImageId != *input.ImageId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	return true
}
