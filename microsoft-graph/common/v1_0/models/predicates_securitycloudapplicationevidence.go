package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidenceOperationPredicate struct {
	AppId                    *int64
	CreatedDateTime          *string
	DisplayName              *string
	InstanceId               *int64
	InstanceName             *string
	ODataType                *string
	RemediationStatusDetails *string
	SaasAppId                *int64
}

func (p SecurityCloudApplicationEvidenceOperationPredicate) Matches(input SecurityCloudApplicationEvidence) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.InstanceId != nil && (input.InstanceId == nil || *p.InstanceId != *input.InstanceId) {
		return false
	}

	if p.InstanceName != nil && (input.InstanceName == nil || *p.InstanceName != *input.InstanceName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediationStatusDetails != nil && (input.RemediationStatusDetails == nil || *p.RemediationStatusDetails != *input.RemediationStatusDetails) {
		return false
	}

	if p.SaasAppId != nil && (input.SaasAppId == nil || *p.SaasAppId != *input.SaasAppId) {
		return false
	}

	return true
}
