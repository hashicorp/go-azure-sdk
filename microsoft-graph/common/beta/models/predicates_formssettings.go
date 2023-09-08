package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormsSettingsOperationPredicate struct {
	IsBingImageSearchEnabled            *bool
	IsExternalSendFormEnabled           *bool
	IsExternalShareCollaborationEnabled *bool
	IsExternalShareResultEnabled        *bool
	IsExternalShareTemplateEnabled      *bool
	IsInOrgFormsPhishingScanEnabled     *bool
	IsRecordIdentityByDefaultEnabled    *bool
	ODataType                           *string
}

func (p FormsSettingsOperationPredicate) Matches(input FormsSettings) bool {

	if p.IsBingImageSearchEnabled != nil && (input.IsBingImageSearchEnabled == nil || *p.IsBingImageSearchEnabled != *input.IsBingImageSearchEnabled) {
		return false
	}

	if p.IsExternalSendFormEnabled != nil && (input.IsExternalSendFormEnabled == nil || *p.IsExternalSendFormEnabled != *input.IsExternalSendFormEnabled) {
		return false
	}

	if p.IsExternalShareCollaborationEnabled != nil && (input.IsExternalShareCollaborationEnabled == nil || *p.IsExternalShareCollaborationEnabled != *input.IsExternalShareCollaborationEnabled) {
		return false
	}

	if p.IsExternalShareResultEnabled != nil && (input.IsExternalShareResultEnabled == nil || *p.IsExternalShareResultEnabled != *input.IsExternalShareResultEnabled) {
		return false
	}

	if p.IsExternalShareTemplateEnabled != nil && (input.IsExternalShareTemplateEnabled == nil || *p.IsExternalShareTemplateEnabled != *input.IsExternalShareTemplateEnabled) {
		return false
	}

	if p.IsInOrgFormsPhishingScanEnabled != nil && (input.IsInOrgFormsPhishingScanEnabled == nil || *p.IsInOrgFormsPhishingScanEnabled != *input.IsInOrgFormsPhishingScanEnabled) {
		return false
	}

	if p.IsRecordIdentityByDefaultEnabled != nil && (input.IsRecordIdentityByDefaultEnabled == nil || *p.IsRecordIdentityByDefaultEnabled != *input.IsRecordIdentityByDefaultEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
