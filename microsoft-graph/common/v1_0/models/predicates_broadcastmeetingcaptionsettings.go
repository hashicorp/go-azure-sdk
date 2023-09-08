package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingCaptionSettingsOperationPredicate struct {
	IsCaptionEnabled *bool
	ODataType        *string
	SpokenLanguage   *string
}

func (p BroadcastMeetingCaptionSettingsOperationPredicate) Matches(input BroadcastMeetingCaptionSettings) bool {

	if p.IsCaptionEnabled != nil && (input.IsCaptionEnabled == nil || *p.IsCaptionEnabled != *input.IsCaptionEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SpokenLanguage != nil && (input.SpokenLanguage == nil || *p.SpokenLanguage != *input.SpokenLanguage) {
		return false
	}

	return true
}
