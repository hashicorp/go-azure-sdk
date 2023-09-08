package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioConferencingOperationPredicate struct {
	ConferenceId   *string
	DialinUrl      *string
	ODataType      *string
	TollFreeNumber *string
	TollNumber     *string
}

func (p AudioConferencingOperationPredicate) Matches(input AudioConferencing) bool {

	if p.ConferenceId != nil && (input.ConferenceId == nil || *p.ConferenceId != *input.ConferenceId) {
		return false
	}

	if p.DialinUrl != nil && (input.DialinUrl == nil || *p.DialinUrl != *input.DialinUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TollFreeNumber != nil && (input.TollFreeNumber == nil || *p.TollFreeNumber != *input.TollFreeNumber) {
		return false
	}

	if p.TollNumber != nil && (input.TollNumber == nil || *p.TollNumber != *input.TollNumber) {
		return false
	}

	return true
}
