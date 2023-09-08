package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidenceOperationPredicate struct {
	ContainerId              *string
	CreatedDateTime          *string
	IsPrivileged             *bool
	Name                     *string
	ODataType                *string
	RemediationStatusDetails *string
}

func (p SecurityContainerEvidenceOperationPredicate) Matches(input SecurityContainerEvidence) bool {

	if p.ContainerId != nil && (input.ContainerId == nil || *p.ContainerId != *input.ContainerId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.IsPrivileged != nil && (input.IsPrivileged == nil || *p.IsPrivileged != *input.IsPrivileged) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
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
