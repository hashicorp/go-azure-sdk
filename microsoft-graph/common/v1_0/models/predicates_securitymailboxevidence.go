package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidenceOperationPredicate struct {
	CreatedDateTime          *string
	DisplayName              *string
	ODataType                *string
	PrimaryAddress           *string
	RemediationStatusDetails *string
}

func (p SecurityMailboxEvidenceOperationPredicate) Matches(input SecurityMailboxEvidence) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrimaryAddress != nil && (input.PrimaryAddress == nil || *p.PrimaryAddress != *input.PrimaryAddress) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	return true
}
