package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOauthApplicationEvidenceOperationPredicate struct {
	AppId                    *string
	CreatedDateTime          *string
	DisplayName              *string
	ODataType                *string
	ObjectId                 *string
	Publisher                *string
	RemediationStatusDetails *string
}

func (p SecurityOauthApplicationEvidenceOperationPredicate) Matches(input SecurityOauthApplicationEvidence) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ObjectId != nil && (input.ObjectId == nil || *p.ObjectId != *input.ObjectId) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	return true
}
