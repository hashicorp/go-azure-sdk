package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptOperationPredicate struct {
	CreatedDateTime          *string
	Description              *string
	DetectionScriptContent   *string
	DisplayName              *string
	EnforceSignatureCheck    *bool
	HighestAvailableVersion  *string
	Id                       *string
	IsGlobalScript           *bool
	LastModifiedDateTime     *string
	ODataType                *string
	Publisher                *string
	RemediationScriptContent *string
	RunAs32Bit               *bool
	Version                  *string
}

func (p DeviceHealthScriptOperationPredicate) Matches(input DeviceHealthScript) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DetectionScriptContent != nil && (input.DetectionScriptContent == nil || *p.DetectionScriptContent != *input.DetectionScriptContent) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnforceSignatureCheck != nil && (input.EnforceSignatureCheck == nil || *p.EnforceSignatureCheck != *input.EnforceSignatureCheck) {
		return false
	}

	if p.HighestAvailableVersion != nil && (input.HighestAvailableVersion == nil || *p.HighestAvailableVersion != *input.HighestAvailableVersion) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsGlobalScript != nil && (input.IsGlobalScript == nil || *p.IsGlobalScript != *input.IsGlobalScript) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.RemediationScriptContent != nil && (input.RemediationScriptContent == nil || *p.RemediationScriptContent != *input.RemediationScriptContent) {
		return false
	}

	if p.RunAs32Bit != nil && (input.RunAs32Bit == nil || *p.RunAs32Bit != *input.RunAs32Bit) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
