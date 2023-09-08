package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceOperationPredicate struct {
	CreatedDateTime          *string
	ODataType                *string
	RemediationStatusDetails *string
	ResourceId               *string
	ResourceName             *string
	ResourceType             *string
}

func (p SecurityAzureResourceEvidenceOperationPredicate) Matches(input SecurityAzureResourceEvidence) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	if p.ResourceName != nil && (input.ResourceName == nil || *p.ResourceName != *input.ResourceName) {
		return false
	}

	if p.ResourceType != nil && (input.ResourceType == nil || *p.ResourceType != *input.ResourceType) {
		return false
	}

	return true
}
