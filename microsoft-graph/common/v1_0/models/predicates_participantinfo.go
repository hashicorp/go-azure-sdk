package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantInfoOperationPredicate struct {
	CountryCode   *string
	LanguageId    *string
	ODataType     *string
	ParticipantId *string
	Region        *string
}

func (p ParticipantInfoOperationPredicate) Matches(input ParticipantInfo) bool {

	if p.CountryCode != nil && (input.CountryCode == nil || *p.CountryCode != *input.CountryCode) {
		return false
	}

	if p.LanguageId != nil && (input.LanguageId == nil || *p.LanguageId != *input.LanguageId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParticipantId != nil && (input.ParticipantId == nil || *p.ParticipantId != *input.ParticipantId) {
		return false
	}

	if p.Region != nil && (input.Region == nil || *p.Region != *input.Region) {
		return false
	}

	return true
}
