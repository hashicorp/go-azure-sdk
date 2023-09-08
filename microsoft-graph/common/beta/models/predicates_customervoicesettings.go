package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerVoiceSettingsOperationPredicate struct {
	IsInOrgFormsPhishingScanEnabled  *bool
	IsRecordIdentityByDefaultEnabled *bool
	IsRestrictedSurveyAccessEnabled  *bool
	ODataType                        *string
}

func (p CustomerVoiceSettingsOperationPredicate) Matches(input CustomerVoiceSettings) bool {

	if p.IsInOrgFormsPhishingScanEnabled != nil && (input.IsInOrgFormsPhishingScanEnabled == nil || *p.IsInOrgFormsPhishingScanEnabled != *input.IsInOrgFormsPhishingScanEnabled) {
		return false
	}

	if p.IsRecordIdentityByDefaultEnabled != nil && (input.IsRecordIdentityByDefaultEnabled == nil || *p.IsRecordIdentityByDefaultEnabled != *input.IsRecordIdentityByDefaultEnabled) {
		return false
	}

	if p.IsRestrictedSurveyAccessEnabled != nil && (input.IsRestrictedSurveyAccessEnabled == nil || *p.IsRestrictedSurveyAccessEnabled != *input.IsRestrictedSurveyAccessEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
